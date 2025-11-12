package integrationtest

import (
	"answerService/internals/app"
	"answerService/internals/config"
	"answerService/internals/logger"
	"answerService/internals/models"
	"answerService/internals/storage/postgres"
	"testing"

	"gorm.io/gorm"
)

func setupTestApp(t *testing.T) *app.App {

	cfg := NewTestConfig()
	logger := logger.GetLogger(cfg.Env)

	application := app.New(cfg, logger)

	db := postgres.GetDBConnect(cfg, logger)
	resetDB(t, db)

	return application
}

func NewTestConfig() *config.Config {
	cfg := &config.Config{
		Env: "DEBUG",
	}
	cfg.App.Port = 8080
	cfg.App.ServerAddr = ":8080"

	cfg.Database.HostLocal = "localhost"
	cfg.Database.HostDocker = "db"
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5432
	cfg.Database.User = "db"
	cfg.Database.Password = "db"
	cfg.Database.Name = "db"

	cfg.Migrations.Path = "./migrations"
	return cfg
}

func resetDB(t *testing.T, db *gorm.DB) {
	// Очистка всех таблиц
	db.Exec("TRUNCATE TABLE answers RESTART IDENTITY CASCADE;")
	db.Exec("TRUNCATE TABLE questions RESTART IDENTITY CASCADE;")

	// Добавление тестового вопроса
	q := models.Question{Text: "What is Go?"}
	if err := db.Create(&q).Error; err != nil {
		t.Fatalf("failed to insert question: %v", err)
	}
}
