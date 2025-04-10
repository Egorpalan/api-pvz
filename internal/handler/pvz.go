package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egorpalan/api-pvz/internal/usecase"
)

type createPVZRequest struct {
	City string `json:"city"`
}

func CreatePVZ(pvzUC *usecase.PVZUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPVZRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		pvz, err := pvzUC.Create(req.City)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pvz)
	}
}
