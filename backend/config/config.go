package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port                 int
	MongoURI             string
	MongoDB              string
	AuthSecret           string
	TokenExpirationHours int `env:"TOKEN_EXPIRATION_HOURS" envDefault:"24"`
}

func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnvOrDefault("PORT", "8080"))
	if err != nil {
		return nil, err
	}

	// Construct MongoDB URI if not provided directly
	mongoURI := getEnvOrDefault("MONGODB_URI", "")
	if mongoURI == "" {
		user := getEnvOrDefault("MONGO_USER", "")
		pass := getEnvOrDefault("MONGO_PASSWORD", "")
		host := getEnvOrDefault("MONGO_HOST", "localhost")
		port := getEnvOrDefault("MONGO_PORT", "27017")
		db := getEnvOrDefault("MONGO_DB", "formease")

		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
			user, pass, host, port, db)
	}

	return &Config{
		Port:                 port,
		MongoURI:             mongoURI,
		MongoDB:              getEnvOrDefault("MONGO_DB", "formease"),
		AuthSecret:           getEnvOrDefault("AUTH_SECRET", ""),
		TokenExpirationHours: getIntEnvOrDefault("JWT_LIFETIME", 24),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnvOrDefault(key string, defaultValue int) int {
	if value, err := strconv.Atoi(getEnvOrDefault(key, "")); err == nil {
		return value
	}
	return defaultValue
}
