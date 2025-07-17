package main

import (
	"log"

	"clean_arc/config"
	"clean_arc/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found or couldn't be loaded. Continuing anyway.")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
