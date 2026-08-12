package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v5/pgxpool"
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

	database := initDB(cfg.DatabaseURL)
	defer database.Close()
}

func initDB(url string) *pgxpool.Pool {
	database, err := db.OpenPool(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v", err)
	}
	return database
}
