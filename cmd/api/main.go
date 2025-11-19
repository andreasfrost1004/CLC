package main

import (
	"log"

	"CLC/internal/config"
	httpserver "CLC/internal/http"
)

func main() {
	cfg := config.Load()

	srv := httpserver.NewServer(cfg)

	log.Printf("API listening on %s\n", cfg.HTTP.Address)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
