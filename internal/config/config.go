package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DATABASE_URL string
	ENVIRONMENT string
	CLERK_API_KEY string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *AppConfig {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &AppConfig{
		DATABASE_URL: getEnv("DATABASE_URL", ""),
		ENVIRONMENT: getEnv("ENVIRONMENT", "development"),
		CLERK_API_KEY: getEnv("CLERK_API_KEY", ""),
	}
}

// getEnv fetches an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
