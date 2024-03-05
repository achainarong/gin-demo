package models

// @Description User model
// Password needs to be hashed but its just a demo, omit is still not working
type User struct {
	Base
	Username string `json:"username" gorm:"type:varchar(100);unique_index" example:"john_doe"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index" example:"john.doe@example.com"`
	Password string `json:"password,omitempty" gorm:"type:varchar(100)" swaggertype:"string" example:"securePassword123"` // omit in JSON output
}
