package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/Egorpalan/api-pvz/internal/usecase"
)

type createReceptionRequest struct {
	PVZID string `json:"pvzId"`
}

func CreateReception(uc *usecase.ReceptionUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createReceptionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		reception, err := uc.Create(req.PVZID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reception)
	}
}

func CloseLastReception(uc *usecase.ReceptionUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pvzID := chi.URLParam(r, "pvzId")
		if pvzID == "" {
			http.Error(w, "missing pvzId", http.StatusBadRequest)
			return
		}

		rec, err := uc.CloseLast(pvzID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rec)
	}
}
