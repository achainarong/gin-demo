package main

import (
	"gin-demo/src/config"
	"gin-demo/src/database"
	// "gin-demo/src/s3"
	"log"
)

func main() {
	settings := config.LoadSettings()

	db, err := database.GetGormDB(settings)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform database migrations
	database.MigrateDB(db)

	// minioClient := s3.NewMinioClient(settings)
}
