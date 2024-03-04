package main

import (
	"gin-demo/src/config"
	"gin-demo/src/database"
	"gin-demo/src/endpoints"
	"github.com/gin-gonic/gin"
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

	// minioClient := s3.NewMinioClient(settings)	// Setup Gin
	r := gin.Default()
	endpoints.RegisterUserRoutes(r, db)
	r.Run()
}
