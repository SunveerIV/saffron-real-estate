package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const mainDatabaseName = "saffron_backend_db"

type Config struct {
	DatabaseURL      string
	MainDatabaseName string
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return Config{}, err
	}
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return Config{}, errors.New("DATABASE_URL environment variable not set")
	}

	return Config{
		DatabaseURL:      url,
		MainDatabaseName: mainDatabaseName,
	}, nil
}
