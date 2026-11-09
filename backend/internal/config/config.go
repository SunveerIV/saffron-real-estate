package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	Port             string
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

	mainDatabaseName := os.Getenv("MAIN_DATABASE_NAME")
	if mainDatabaseName == "" {
		return Config{}, errors.New("MAIN_DATABASE_NAME environment variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return Config{}, errors.New("PORT environment variable not set")
	}

	return Config{
		DatabaseURL:      url,
		MainDatabaseName: mainDatabaseName,
		Port:             port,
	}, nil
}
