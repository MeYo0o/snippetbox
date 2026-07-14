package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func setupDB(logger *slog.Logger) (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Couldn't read .env")
		return nil, err
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logger.Error("Couldn't Get DBURl")
		return nil, err
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error(), "DB Connection issue", "Couldn't connect to DB")
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
