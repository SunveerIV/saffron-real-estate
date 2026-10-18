package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
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
		DatabaseURL: url,
	}, nil
}
