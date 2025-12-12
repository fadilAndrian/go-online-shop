package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBDSN   string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("ENV not found")
	}

	cfg := Config{
		AppPort: os.Getenv("APP_PORT"),
		DBDSN:   os.Getenv("DB_DSN"),
	}

	if cfg.AppPort == "" || cfg.DBDSN == "" {
		log.Fatal("Missing environment variables")
	}

	return cfg
}
