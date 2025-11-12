package migrator

import (
	"answerService/internals/config"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func RunMigrations(cfg *config.Config, logger *slog.Logger) error {

	if cfg.Database.User == "" ||
		cfg.Database.Password == "" ||
		cfg.Database.Port == 0 ||
		cfg.Database.Name == "" ||
		cfg.Database.Host == "" {
		return fmt.Errorf(
			"incomplete DB configuration: user=%q, name=%q, host=%q, port=%d",
			cfg.Database.User, cfg.Database.Name, cfg.Database.Host, cfg.Database.Port,
		)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name,
	)

	logger.Info("Starting database migrations")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer func() {
		if cerr := db.Close(); cerr != nil {
			logger.Warn("failed to close DB connection", slog.String("error", cerr.Error()))
		}
	}()

	migrationPath := cfg.Migrations.Path

	logger.Info("Using migrations path", slog.String("path", migrationPath))

	goose.SetBaseFS(nil)
	goose.SetLogger(goose.NopLogger())

	if err := goose.Up(db, migrationPath); err != nil {
		if err.Error() == "no migrations to run. current version: 1" {
			logger.Info("No new migrations to apply")
		} else {
			return fmt.Errorf("run migrations: %w", err)
		}
	}

	logger.Info("Database migrations ran successfully")
	return nil
}

func MustRunMigrations(cfg *config.Config, logger *slog.Logger) {
	err := RunMigrations(cfg, logger)
	if err != nil {
		panic(err)
	}
}
