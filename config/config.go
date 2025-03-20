package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Default config values
const (
	// Postgres
	POSTGRES_HOST     = "localhost"
	POSTGRES_PORT     = "5432"
	POSTGRES_USER     = "postgres"
	POSTGRES_PASSWORD = "postgres"
	POSTGRES_DB       = "postgres"
)

type Config struct {
	PostgresConfig postgresConfig
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	if os.Getenv("POSTGRES_HOST") == "" {
		os.Setenv("POSTGRES_HOST", POSTGRES_HOST)
	}
	if os.Getenv("POSTGRES_PORT") == "" {
		os.Setenv("POSTGRES_PORT", POSTGRES_PORT)
	}
	if os.Getenv("POSTGRES_USER") == "" {
		os.Setenv("POSTGRES_USER", POSTGRES_USER)
	}
	if os.Getenv("POSTGRES_PASSWORD") == "" {
		os.Setenv("POSTGRES_PASSWORD", POSTGRES_PASSWORD)
	}
	if os.Getenv("POSTGRES_DB") == "" {
		os.Setenv("POSTGRES_DB", POSTGRES_DB)
	}

	return Config{}
}

func (c Config) WithPostgres() Config {
	c.PostgresConfig = postgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	return c
}
