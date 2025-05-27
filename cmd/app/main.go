package main

import (
	"log"
	"os"

	"github.com/andreyxaxa/Quotation-Book-service/config"
	"github.com/andreyxaxa/Quotation-Book-service/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("No .env file found")
		}
	}

	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
