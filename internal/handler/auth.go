package handler

import (
	"encoding/json"
	"github.com/Egorpalan/api-pvz/internal/usecase"
	"net/http"

	"github.com/Egorpalan/api-pvz/pkg/jwt"
)

type DummyLoginRequest struct {
	Role string `json:"role"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	var req DummyLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || (req.Role != "employee" && req.Role != "moderator") {
		http.Error(w, "invalid role", http.StatusBadRequest)
		return
	}

	token, err := jwt.GenerateToken(req.Role)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := TokenResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Register(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		user, err := uc.Register(req.Email, req.Password, req.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func Login(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		token, err := uc.Login(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
