// Package models
package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	Email     string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Age       int       `gorm:"default:18" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// 	// One-to-Many relationship
	// 	Posts []Post `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

// 	// Relationships
// 	Posts []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
// }

// // TableName specifies the table name for User
// func (User) TableName() string {
// 	return "users"
// }

// // BeforeCreate is a GORM hook that runs before creating a user
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	// You can add validation logic here
// 	if u.Age < 0 {
// 		u.Age = 0
// 	}
// 	return nil
// }

// // GetFullName returns the user's full name
// func (u *User) GetFullName() string {
// 	return u.Name
// }

// // IsAdult returns true if the user is 18 or older
// func (u *User) IsAdult() bool {
// 	return u.Age >= 18
// }
