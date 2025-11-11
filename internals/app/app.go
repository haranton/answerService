package app

import (
	"authservice/internals/config"
	"authservice/internals/handlers"
	"authservice/internals/service"
	"authservice/internals/storage"
	"authservice/internals/storage/migrator"
	"authservice/internals/storage/postgres"
	"log/slog"
	"net/http"
)

type App struct {
	cfg     *config.Config
	logger  *slog.Logger
	storage storage.Storage
	service *service.Service
	server  *handlers.Handler
}

func New(cfg *config.Config, logger *slog.Logger) *App {

	dbConnect := postgres.GetDBConnect(cfg, logger)

	st := postgres.NewPostgresStorage(dbConnect)

	svc := service.NewService(st, logger, cfg)

	srv := handlers.NewHandler(svc, logger, cfg)

	return &App{
		cfg:     cfg,
		logger:  logger,
		storage: st,
		service: svc,
		server:  srv,
	}
}

func (a *App) MustStart() error {

	migrator.MustRunMigrations(a.cfg, a.logger)

	mux := http.NewServeMux()
	a.server.RegisterRoutes(mux)

	a.logger.Info("starting HTTP server", slog.String("addr", a.cfg.App.ServerAddr))
	return http.ListenAndServe(a.cfg.App.ServerAddr, mux)
}

func (a *App) Close() {
	a.storage.Close()
}
