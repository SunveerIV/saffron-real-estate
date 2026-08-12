package main

import (
	"log"

	_ "github.com/lib/pq"

	"saffronreal.estate/backend/internal/config"
	"saffronreal.estate/backend/internal/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	err = db.RunMigrations(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	_, err = db.OpenPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v", err)
	}
}
