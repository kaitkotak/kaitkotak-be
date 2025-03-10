package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds database configuration
type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
}

// LoadConfig reads configuration from `configs/.env`
func LoadConfig() *Config {
	envPath := filepath.Join("configs", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("No .env file found at %s, using system environment variables\n", envPath)
	}

	return &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
	}
}
