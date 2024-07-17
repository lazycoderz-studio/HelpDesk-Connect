package db

import (
	"lazycoderz-studio/helpdesk-connect/connect/models"
	"log"
)

// Migrate performs the database migrations
func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
