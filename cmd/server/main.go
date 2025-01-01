package main

import (
    "fmt"
    "log"

    "github.com/twjsanderson/decision_backend/api/routes"
	"github.com/twjsanderson/decision_backend/internal/config"
)

func main() {
	// load config
	appConfig := config.LoadConfig()

	fmt.Println("DB URL", appConfig.DatabaseURL)

	// initialize server routes
	r := routes.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}
