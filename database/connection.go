package database

import (
	"log"

	"go_habbit/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("habbit.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

// Migrate runs database migrations
func Migrate() {
	// Import your models here
	err := DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitCompletion{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
