package db

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"saffronreal.estate/backend/internal/util"
)

func RunMigrations(pool *pgxpool.Pool) error {
	conn := stdlib.OpenDBFromPool(pool)
	defer conn.Close()

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create database driver: %v", err)
	}

	backendDirectory, err := util.FindBackendDirectory()
	if err != nil {
		return fmt.Errorf("unable to find backend directory: %w", err)
	}

	sourceURL := url.URL{
		Scheme: "file",
		Path:   filepath.Join(backendDirectory, "migrations"),
	}

	m, err := migrate.NewWithDatabaseInstance(
		sourceURL.String(),
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("unable to create migration instance: %v", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to run migrations: %v", err)
	}

	return nil
}
