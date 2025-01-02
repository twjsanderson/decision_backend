package main

import (
	"log"

	"github.com/twjsanderson/decision_backend/api/routes"
	// "github.com/twjsanderson/decision_backend/internal/config"
)

func main() {
	// load config
	// appConfig := config.LoadConfig()

	// fmt.Println("DB URL", appConfig.DATABASE)URL)

	// initialize server routes
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
