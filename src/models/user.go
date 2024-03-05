package models

import (
	"github.com/google/uuid"
)

// @Description User model
// Password needs to be hashed but its just a demo
type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username string    `json:"username" gorm:"type:varchar(100);unique_index" example:"john_doe"` // Swagger example
	Email    string    `json:"email" gorm:"type:varchar(100);unique_index" example:"john.doe@example.com"`
	Password string    `json:"password,omitempty" gorm:"type:varchar(100)" swaggertype:"string" example:"securePassword123"` // omit in JSON output
}
