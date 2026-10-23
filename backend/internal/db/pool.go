package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func PoolWithName(
	ctx context.Context,
	adminPool *pgxpool.Pool,
	name string,
) (*pgxpool.Pool, error) {
	var exists bool

	err := adminPool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
		name,
	).Scan(&exists)

	if !exists {
		_, err := adminPool.Exec(ctx, "CREATE DATABASE "+pgx.Identifier{name}.Sanitize())
		if err != nil {
			return nil, fmt.Errorf("error creating main database: %w", err)
		}
	}

	cfg := adminPool.Config()
	cfg.ConnConfig.Database = name

	mainDatabase, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to main database: %w", err)
	}
	return mainDatabase, nil
}
