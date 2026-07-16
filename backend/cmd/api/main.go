package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db := initDB()
	defer db.Close()
}

func initDB() *pgxpool.Pool {
	db, err := pgxpool.New(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return db
}
