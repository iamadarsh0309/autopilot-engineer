package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Using port %s", port)
	return &Config{Port: port}
}
