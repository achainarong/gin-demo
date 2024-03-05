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

	// Needs to be created to generate uuids
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Perform database migrations
	database.MigrateDB(db)

	// minioClient := s3.NewMinioClient(settings)	// Setup Gin
	r := gin.Default()
	endpoints.RegisterUserRoutes(r, db)
	r.Run()
}
