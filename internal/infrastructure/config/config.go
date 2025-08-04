package config

import (
	"chekisvc/internal/infrastructure/database"
	"log"
	"os"
)

// Config represents application configuration
type Config struct {
	Server   ServerConfig
	Database database.Config
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: database.Config{
			Host:     getEnv("DB_HOST", "172.18.125.255"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "vini"),
			DBName:   getEnv("DB_NAME", "chekisvc"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
	}
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		log.Printf("[CONFIG] %s loaded from ENV: %s", key, value)
		return value
	}
	log.Printf("[CONFIG] %s using default: %s", key, defaultValue)
	return defaultValue
}
