package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the service
type Config struct {
	// Service info
	ServiceName string
	Version     string
	Environment string

	// Server
	Port string

	// External services
	ItemServiceURL    string
	PaymentServiceURL string

	// CORS
	CORSOrigin string
}

// Load reads configuration from environment variables
func Load() *Config {
	// Load .env file if it exists (ignore error in production)
	_ = godotenv.Load()

	return &Config{
		ServiceName:       getEnv("SERVICE_NAME", "order-service"),
		Version:           getEnv("SERVICE_VERSION", "1.0.0"),
		Environment:       getEnv("GO_ENV", "development"),
		Port:              getEnv("PORT", "3002"),
		ItemServiceURL:    getEnv("ITEM_SERVICE_URL", "http://localhost:3001"),
		PaymentServiceURL: getEnv("PAYMENT_SERVICE_URL", "http://localhost:3003"),
		CORSOrigin:        getEnv("CORS_ORIGIN", "*"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
