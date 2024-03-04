package models

import (
	"gorm.io/gorm"
)

// User defines the model for a user with GORM annotations for the database schema
type User struct {
	gorm.Model        // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Username  string `gorm:"type:varchar(100);unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string `gorm:"type:varchar(100)"`
}

