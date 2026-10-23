package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"

	"saffronreal.estate/backend/internal/config"
	"saffronreal.estate/backend/internal/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	pgxcfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	adminPool, err := pgxpool.NewWithConfig(ctx, pgxcfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer adminPool.Close()

	mainPool, err := db.PoolWithName(ctx, adminPool, cfg.MainDatabaseName)
	if err != nil {
		log.Fatalf("Unable to retrieve database main database: %v", err)
	}
	defer mainPool.Close()

	err = db.RunMigrations(mainPool)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}
