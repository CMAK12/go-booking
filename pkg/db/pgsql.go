package db

import (
	"context"
	"fmt"
	"go-booking/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustConnect(ctx context.Context, pgConfig config.PostgresConfig) (*pgxpool.Pool, string) {
	dsn := pgConfig.BuildConnectionString()

	poolConfig, err := pgxpool.ParseConfig(dsn)
	fmt.Println("DSN:", dsn)
	if err != nil {
		log.Fatalf("failed to parse database URL: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("database connection established")

	return pool, dsn
}
