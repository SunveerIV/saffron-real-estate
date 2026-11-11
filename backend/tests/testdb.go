package tests

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"saffronreal.estate/backend/internal/config"
	"saffronreal.estate/backend/internal/db"
)

var ctx context.Context
var adminPool *pgxpool.Pool

func init() {
	err := run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	ctx = context.Background()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("unable to load config: %v", err)
	}

	pgxcfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("unable to parse config: %v", err)
	}

	adminPool, err = pgxpool.NewWithConfig(ctx, pgxcfg)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	return nil
}

func newDatabaseName() string {
	newUUID := uuid.New()
	return "test_" + strings.ReplaceAll(newUUID.String(), "-", "")
}

func NewTestDB() *pgxpool.Pool {
	testPool, err := db.PoolWithName(ctx, adminPool, newDatabaseName())
	if err != nil {
		log.Fatalf("unable to retrieve database main database: %v", err)
	}

	err = db.RunMigrations(testPool)
	if err != nil {
		log.Fatalf("unable to migrate database: %v", err)
	}
	return testPool
}

func DeleteTestDB(pool *pgxpool.Pool) {
	name := pool.Config().ConnConfig.Database
	pool.Close()

	query := "DROP DATABASE " + pgx.Identifier{name}.Sanitize()
	_, err := adminPool.Exec(ctx, query)
	if err != nil {
		log.Fatalf("unable to delete test database: %v", err)
	}
}
