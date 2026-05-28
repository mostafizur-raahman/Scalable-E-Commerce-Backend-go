package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string, maxConns int32) (*DB, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse db config: %w", err)
	}

	config.MaxConns = maxConns
	config.MinConns = 2
	config.MaxConnIdleTime = 30 * time.Minute
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return &DB{Pool: pool}, nil
}

func InitSchema(ctx context.Context, db *DB, schemaPath string) error {
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("read schema file: %w", err)
	}

	_, err = db.Pool.Exec(ctx, string(schema))
	if err != nil {
		return fmt.Errorf("execute schema: %w", err)
	}

	slog.Info("✅ Schema initialized from file", "path", schemaPath)
	return nil
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		slog.Info("🔌 Database disconnected")
	}
}
