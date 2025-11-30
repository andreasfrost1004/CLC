package config

import (
	"os"
)

type HTTPConfig struct {
	Address string
}

// for blizzard API for items
type BlizzardConfig struct {
	ClientID     string
	ClientSecret string
	Region       string
	Namespace    string
	Locale       string
}

type Config struct {
	HTTP     HTTPConfig
	Blizzard BlizzardConfig
}

func Load() Config {
	addr := os.Getenv("API_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return Config{
		HTTP: HTTPConfig{
			Address: addr,
		},
		Blizzard: BlizzardConfig{
			ClientID:     os.Getenv("BLIZZARD_CLIENT_ID"),
			ClientSecret: os.Getenv("BLIZZARD_CLIENT_SECRET"),
			Region:       os.Getenv("BLIZZARD_REGION"),
			Namespace:    os.Getenv("BLIZZARD_NAMESPACE"),
			Locale:       os.Getenv("BLIZZARD_LOCALE"),
		},
	}
}
