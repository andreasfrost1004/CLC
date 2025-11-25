package main

import (
	"log"

	"CLC/internal/config"
	"CLC/internal/database"
	httpserver "CLC/internal/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // loads .env into os.Environ

	cfg := config.Load()

	db, err := database.New()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	srv := httpserver.NewServer(cfg, db)

	log.Printf("API listening on %s\n", cfg.HTTP.Address)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
