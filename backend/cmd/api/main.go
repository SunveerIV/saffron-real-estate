package main

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/config"
	"saffronreal.estate/backend/internal/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var url string = initURL()
	runMigrations(url)
	database := initDB(url)
	defer database.Close()
}

func initURL() string {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to get env DB URL: %v", err)
	}
	return cfg.DatabaseURL
}

func initDB(url string) *pgxpool.Pool {
	database, err := db.OpenPool(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v", err)
	}
	return database
}

func runMigrations(databaseURL string) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close()

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Unable to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///backend/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Unable to create migration instance: %v", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Unable to run migrations: %v", err)
	}
}
