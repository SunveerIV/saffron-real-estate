package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/app"

	_ "github.com/lib/pq"

	"saffronreal.estate/backend/internal/config"
	"saffronreal.estate/backend/internal/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("unable to load config: %v", err)
	}

	pgxcfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("unable to parse config: %v", err)
	}

	adminPool, err := pgxpool.NewWithConfig(ctx, pgxcfg)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	defer adminPool.Close()

	mainPool, err := db.PoolWithName(ctx, adminPool, cfg.MainDatabaseName)
	if err != nil {
		return fmt.Errorf("unable to retrieve database main database: %v", err)
	}
	defer mainPool.Close()

	err = db.RunMigrations(mainPool)
	if err != nil {
		return fmt.Errorf("unable to migrate database: %v", err)
	}

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: app.New(mainPool),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("HTTP server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	return <-serverErrors
}
