package config

// ###############################################################################
// Configurations
// ###############################################################################

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	Version       string
	Port          string
	Environment   string
	Description   string
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() (*Config, error) {
	godotenv.Load() // Ignore error if .env file doesn't exist

	return &Config{
		AppName:       getEnv("APP_NAME", "Naqa API"),
		Version:       getEnv("API_VERSION", "1.0.0"),
		Port:          getEnv("PORT", "3000"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		Description:   getEnv("APP_DESCRIPTION", ""),
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase: getEnv("MONGO_DATABASE", "naqa"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
