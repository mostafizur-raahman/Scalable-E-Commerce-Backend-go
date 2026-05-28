package database

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations executes pending migrations against the database
func RunMigrations(dsn, migrationsDir string) error {
	// Convert to absolute path to avoid CWD issues
	absDir, err := filepath.Abs(migrationsDir)
	if err != nil {
		return fmt.Errorf("resolve migrations path: %w", err)
	}

	slog.Info("🔄 Running database migrations...", "path", absDir)

	// Create migrate instance
	m, err := migrate.New(
		fmt.Sprintf("file://%s", absDir),
		dsn,
	)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}

	// Run all pending migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("✅ No pending migrations")
			return nil
		}
		return fmt.Errorf("apply migrations: %w", err)
	}

	slog.Info("✅ Database migrations completed successfully")
	return nil
}

// RollbackMigrations rolls back the last migration (for development)
func RollbackMigrations(dsn, migrationsDir string) error {
	absDir, err := filepath.Abs(migrationsDir)
	if err != nil {
		return fmt.Errorf("resolve migrations path: %w", err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", absDir),
		dsn,
	)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("rollback migration: %w", err)
	}

	slog.Info("✅ Migration rolled back successfully")
	return nil
}
