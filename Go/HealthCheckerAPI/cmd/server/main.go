package main

import (
	"HealthCheckerAPI/pkg/db"
	"HealthCheckerAPI/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}

	db.RunMigrations()

	http.HandleFunc("/status", handlers.GetHealthStatus)
	http.HandleFunc("/status/", handlers.GetHealthStatusByServerID)
	http.HandleFunc("/healthcheck", handlers.PostHealthStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf(".env file not found or couldn't be loaded: %w", err)
	}
	return nil
}
