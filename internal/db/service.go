package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/twjsanderson/decision_backend/internal/config"
)

var DB *pgxpool.Pool

// InitializeDB initializes a connection pool to the PostgreSQL database
func InitializeDB() error {
	appConfig := config.LoadConfig()
	dbURL := appConfig.DATABASE_URL
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	// Configure the connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("unable to parse database URL: %w", err)
	}

	// Create a new connection pool
	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %w", err)
	}

	// Test the connection
	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("database connection test failed: %w", err)
	}

	log.Println("Connected to the database successfully")
	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
