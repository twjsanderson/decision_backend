package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds the application's configuration
type AppConfig struct {
	DatabaseURL string
	Environment string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *AppConfig {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &AppConfig{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://default_url"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv fetches an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
