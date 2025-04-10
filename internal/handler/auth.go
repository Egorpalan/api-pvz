package handler

import (
	"encoding/json"
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || (req.Role != "client" && req.Role != "moderator") {
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
