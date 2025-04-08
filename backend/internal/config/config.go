package config

import (
	"os"
)

type Config struct {
	DBPath     string
	JWTSecret  string
	ServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBPath:     getEnv("DB_PATH", "wallos.db"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 