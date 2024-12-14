package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
}

func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnvOrDefault("PORT", "8080"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:        port,
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/formease?sslmode=disable"),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
