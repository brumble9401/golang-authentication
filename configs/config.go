package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds the configuration values (e.g., database URL).
type Config struct {
	DatabaseUrl string
}

var Envs = initConfig()

// initConfig loads the environment variables and returns the configuration.
func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config {
		DatabaseUrl: getEnv("DATABASE_URL"),
	}
}

// getEnv retrieves an environment variable or throws an error if it doesn't exist.
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set.", key)
	}
	return value
}