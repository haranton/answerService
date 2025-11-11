package handlers

import (
	"authservice/internals/dto"
	"authservice/internals/models"
	"authservice/internals/storage/postgres"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

func (h *Handler) UsersRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		h.logger.Error("failed to decode user", slog.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if reqUser.Telephone == "" {
		writeError(w, http.StatusBadRequest, "Telephone is required")
		return
	}
	if reqUser.Password == "" {
		writeError(w, http.StatusBadRequest, "Password is required")
		return
	}

	user := models.User{
		Telephone:    reqUser.Telephone,
		PasswordHash: reqUser.Password,
	}

	registeredUser, accessToken, rawRefresh, err := h.service.SrvUser.RegisterUser(r.Context(), &user)
	if err != nil {
		h.logger.Error("failed to register user", slog.String("error", err.Error()))

		if errors.Is(err, postgres.ErrUserAlreadyExists) {
			writeError(w, http.StatusConflict, "User with this telephone already exists")
			return
		}

		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp := dto.UserResponse{
		UserID:    registeredUser.ID,
		Telephone: registeredUser.Telephone,
		CreatedAt: registeredUser.CreatedAt.Format(time.RFC3339),
		UserLoginResponse: dto.UserLoginResponse{
			AccessToken:  accessToken,
			RefreshToken: rawRefresh,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("failed to write response", slog.String("error", err.Error()))
	}
}

func (h *Handler) UsersLoginHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		h.logger.Error("failed to decode user", slog.String("error", err.Error()))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if reqUser.Telephone == "" {
		writeError(w, http.StatusBadRequest, "Telephone is required")
		return
	}
	if reqUser.Password == "" {
		writeError(w, http.StatusBadRequest, "Password is required")
		return
	}

	accessToken, refreshToken, err := h.service.SrvUser.LoginUser(r.Context(), reqUser.Telephone, reqUser.Password)
	if err != nil {
		h.logger.Error("failed to login user", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) UsersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.logger.Error("failed to decode logout body", slog.String("error", err.Error()))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.SrvUser.Logout(r.Context(), body.RefreshToken); err != nil {
		h.logger.Error("failed to logout user", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "logged out",
	})
}
