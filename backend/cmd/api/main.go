package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/config"
)

func main() {
	db := initDB()
	defer db.Close()
}

func initDB() *pgxpool.Pool {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to get env DB URL: %v", err)
	}

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return db
}
