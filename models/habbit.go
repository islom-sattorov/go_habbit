package models

import "time"

type Habit struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"size:500" json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	StartDate   time.Time `gorm:"not null" json:"start_date"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`

	// Cached streak values for performance
	CurrentStreak   int        `gorm:"default:0" json:"current_streak"`
	LongestStreak   int        `gorm:"default:0" json:"longest_streak"`
	LastCompletedAt *time.Time `json:"last_completed_at"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	User        User              `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Completions []HabitCompletion `gorm:"foreignKey:HabitID" json:"completions,omitempty"`
}

func (Habit) TableName() string {
	return "habits"
}
