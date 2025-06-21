package services

import (
	"errors"
	"time"

	"go_habbit/models"

	"gorm.io/gorm"
)

type HabitService struct {
	db *gorm.DB
}

func NewHabitService(db *gorm.DB) *HabitService {
	return &HabitService{db: db}
}

// CompleteHabit marks a habit as completed for today
func (s *HabitService) CompleteHabit(userID, habitID uint) error {
	// Check if already completed today
	today := time.Now().Truncate(24 * time.Hour)

	var existingCompletion models.HabitCompletion
	err := s.db.Where("habit_id = ? AND user_id = ? AND DATE(completed_at) = DATE(?)",
		habitID, userID, today).First(&existingCompletion).Error

	if err == nil {
		return errors.New("habit already completed today")
	}

	// Create completion record
	completion := models.HabitCompletion{
		HabitID:     habitID,
		UserID:      userID,
		CompletedAt: time.Now(),
	}

	if err := s.db.Create(&completion).Error; err != nil {
		return err
	}

	// Update streak counters
	return s.updateStreaks(habitID)
}

// updateStreaks recalculates and updates streak counters
func (s *HabitService) updateStreaks(habitID uint) error {
	var habit models.Habit
	if err := s.db.First(&habit, habitID).Error; err != nil {
		return err
	}

	// Calculate current streak
	currentStreak := s.calculateCurrentStreak(habitID)

	// Update longest streak if current is longer
	longestStreak := habit.LongestStreak
	if currentStreak > longestStreak {
		longestStreak = currentStreak
	}

	// Update habit with new streaks
	now := time.Now()
	return s.db.Model(&habit).Updates(map[string]interface{}{
		"current_streak":    currentStreak,
		"longest_streak":    longestStreak,
		"last_completed_at": &now,
	}).Error
}

// calculateCurrentStreak calculates consecutive days from today backwards
func (s *HabitService) calculateCurrentStreak(habitID uint) int {
	var completions []models.HabitCompletion

	// Get completions ordered by date desc
	s.db.Where("habit_id = ?", habitID).
		Order("completed_at DESC").
		Find(&completions)

	if len(completions) == 0 {
		return 0
	}

	streak := 0
	today := time.Now().Truncate(24 * time.Hour)

	// Check if completed today or yesterday (grace period)
	lastCompletion := completions[0].CompletedAt.Truncate(24 * time.Hour)
	daysDiff := int(today.Sub(lastCompletion).Hours() / 24)

	if daysDiff > 1 {
		return 0 // Streak broken
	}

	// Count consecutive days
	expectedDate := lastCompletion
	for _, completion := range completions {
		completionDate := completion.CompletedAt.Truncate(24 * time.Hour)
		if completionDate.Equal(expectedDate) {
			streak++
			expectedDate = expectedDate.AddDate(0, 0, -1)
		} else {
			break
		}
	}

	return streak
}

// IsCompletedToday checks if habit is completed today
func (s *HabitService) IsCompletedToday(userID, habitID uint) (bool, error) {
	today := time.Now().Truncate(24 * time.Hour)

	var count int64
	err := s.db.Model(&models.HabitCompletion{}).
		Where("habit_id = ? AND user_id = ? AND DATE(completed_at) = DATE(?)",
			habitID, userID, today).
		Count(&count).Error

	return count > 0, err
}

// GetHabitStats returns comprehensive habit statistics
func (s *HabitService) GetHabitStats(habitID uint) (*HabitStats, error) {
	var habit models.Habit
	if err := s.db.First(&habit, habitID).Error; err != nil {
		return nil, err
	}

	var totalCompletions int64
	s.db.Model(&models.HabitCompletion{}).Where("habit_id = ?", habitID).Count(&totalCompletions)

	// Calculate completion rate (last 30 days)
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	var completionsLast30Days int64
	s.db.Model(&models.HabitCompletion{}).
		Where("habit_id = ? AND completed_at >= ?", habitID, thirtyDaysAgo).
		Count(&completionsLast30Days)

	completionRate := float64(completionsLast30Days) / 30.0 * 100

	return &HabitStats{
		HabitID:           habit.ID,
		CurrentStreak:     habit.CurrentStreak,
		LongestStreak:     habit.LongestStreak,
		TotalCompletions:  int(totalCompletions),
		CompletionRate30D: completionRate,
		LastCompletedAt:   habit.LastCompletedAt,
	}, nil
}

type HabitStats struct {
	HabitID           uint       `json:"habit_id"`
	CurrentStreak     int        `json:"current_streak"`
	LongestStreak     int        `json:"longest_streak"`
	TotalCompletions  int        `json:"total_completions"`
	CompletionRate30D float64    `json:"completion_rate_30d"`
	LastCompletedAt   *time.Time `json:"last_completed_at"`
}
