package postgres

import (
	"fmt"
	"log/slog"
	"os"

	"answerService/internals/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnect(config *config.Config, logger *slog.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	logger.Info("Connecting to database", slog.String("dsn", dsn))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return db
}
