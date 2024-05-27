package store

import (
	"github.com/ArenDjango/golang-test-task/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runPgMigrations() error {
	cfg := config.Get()
	if cfg.PgMigrationsPath == "" {
		return errors.New("PgMigrationsPath is empty")
	}
	if cfg.PgURL == "" {
		return errors.New("No cfg.PgURL provided")
	}

	// Check if migration path exists
	if _, err := os.Stat(cfg.PgMigrationsPath); os.IsNotExist(err) {
		return errors.New("PgMigrationsPath does not exist")
	}

	m, err := migrate.New(
		"file://"+cfg.PgMigrationsPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
