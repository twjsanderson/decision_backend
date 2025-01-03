package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/twjsanderson/decision_backend/api/routes"
	"github.com/twjsanderson/decision_backend/internal/db"
)

func main() {
	// Initialize the database
	if err := db.InitializeDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Setup routes
	router := routes.SetupRouter()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		db.CloseDB()
	}()

	// Start the server
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
