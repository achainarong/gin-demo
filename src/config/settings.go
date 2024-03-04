package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	Host              string `envconfig:"HOST"`
	Port              int    `envconfig:"PORT"`
	PostgresHost      string `envconfig:"POSTGRES_HOST"`
	PostgresDBName    string `envconfig:"POSTGRES_DB_NAME"`
	PostgresPassword  string `envconfig:"POSTGRES_PASSWORD"`
	PostgresUser      string `envconfig:"POSTGRES_USER"`
	PostgresPort      int    `envconfig:"POSTGRES_PORT"`
	MinioAccessKey    string `envconfig:"MINIO_ACCESS_KEY"`
	MinioAccessSecret string `envconfig:"MINIO_ACCESS_SECRET"`
	MinioHost         string `envconfig:"MINIO_HOST"`
}

func LoadSettings() *Settings {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	var settings Settings
	if err := envconfig.Process("", &settings); err != nil {
		log.Fatalf("Error processing environment variables: %s", err)
	}

	return &settings
}
