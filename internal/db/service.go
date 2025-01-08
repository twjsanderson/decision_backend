package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/twjsanderson/decision_backend/internal/config"
	"github.com/twjsanderson/decision_backend/internal/models"
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

	if err := SetupDB(); err != nil {
		return fmt.Errorf("database tables creation failed: %w", err)
	}

	log.Println("Connected to the database successfully")
	return nil
}

func SetupDB() error {
	tables := []models.NewTable{
		{
			TableName:  "Users",
			PrimaryKey: "id",
			Columns: map[string]string{
				"id":         "TEXT",
				"email":      "TEXT",
				"first_name": "TEXT",
				"last_name":  "TEXT",
			},
		},
	}

	for _, table := range tables {
		err := CreateTable(table.TableName, table.Columns)
		if err != nil {
			return fmt.Errorf("error in CreateTable - %v: %v", table.TableName, err)
		}
	}
	return nil
}

func CreateTable(tableName string, columns map[string]string) error {
	// Validate inputs
	if tableName == "" || len(columns) == 0 {
		return fmt.Errorf("table name and columns are required")
	}

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)

	for colName, colType := range columns {
		// Check if this column is marked as PRIMARY KEY
		if colType == "PRIMARY KEY" {
			query += fmt.Sprintf("%s TEXT %s, ", colName, colType)
		}
		query += fmt.Sprintf("%s %s, ", colName, colType)
	}

	// Remove the trailing comma and space
	query = query[:len(query)-2] + ");"

	// Print the final query for debugging purposes
	fmt.Printf("Executing Query: %s\n", query)

	// Execute the query
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	fmt.Printf("Table '%s' created successfully (or already exists)\n", tableName)
	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
