package config

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseURL string
}

func Load() (Config, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return Config{}, errors.New("DATABASE_URL environment variable not set")
	}

	return Config{
		DatabaseURL: url,
	}, nil
}
