package handlers

import (
	"authservice/internals/config"
	"authservice/internals/service"
	"log/slog"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  *slog.Logger
	cfg     *config.Config
}

func NewHandler(service *service.Service, logger *slog.Logger, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/auth/register", h.UsersRegisterHandler)
	mux.HandleFunc("POST /api/auth/login", h.UsersLoginHandler)
	mux.HandleFunc("POST /api/auth/logout", h.UsersLogoutHandler)
	mux.HandleFunc("POST /api/auth/refresh", h.UsersRefreshHandler)

	//todo proxy
	mux.Handle("/api/users/", h.AuthProxyHandler(h.cfg.ServiceNginxURL))

}
