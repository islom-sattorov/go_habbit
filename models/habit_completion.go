package models

import (
	"time"
)

type HabitCompletion struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	HabitID     uint      `gorm:"not null;index" json:"habit_id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	CompletedAt time.Time `gorm:"not null" json:"completed_at"`
	Notes       string    `gorm:"size:500" json:"notes,omitempty"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Habit Habit `gorm:"foreignKey:HabitID" json:"habit,omitempty"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (HabitCompletion) TableName() string {
	return "habit_completions"
}
