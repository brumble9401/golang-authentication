package db

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func EnvVar(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			return defaultValue
		}
		log.Fatalf("Environment variable %s is required but not set.", key)
	}
	return value
}

func NewPostgresStorage(dbUrl string) (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Error("Error connect to Postgres database: %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Error("Error pinging database: %v", err)
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Error("Error pinging database: %v", err)
		return nil, err
	}

	log.Error("Successfully connected to the PostgreSQL database!")
	return pool, nil
}