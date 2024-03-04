package database

import (
	"fmt"
	"gin-demo/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormDB(settings *config.Settings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		settings.PostgresHost, settings.PostgresUser, settings.PostgresPassword, settings.PostgresDBName, settings.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
