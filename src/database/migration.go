package database

import (
	"gin-demo/src/models" 
	"gorm.io/gorm"
	"log"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		// Add other models here, e.g., &models.Product{}, &models.Order{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}
}

