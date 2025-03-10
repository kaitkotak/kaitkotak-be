package config

import (
	"os"
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
	return &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
	}
}
