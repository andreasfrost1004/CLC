package config

import (
	"os"
)

type HTTPConfig struct {
	Address string
}

type Config struct {
	HTTP HTTPConfig
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
	}
}
