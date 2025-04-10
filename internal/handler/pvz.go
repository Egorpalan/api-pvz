package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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

func GetPVZList(pvzUC *usecase.PVZUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			startDate, endDate *time.Time
			page               = 1
			limit              = 10
		)

		if s := r.URL.Query().Get("startDate"); s != "" {
			t, err := time.Parse(time.RFC3339, s)
			if err == nil {
				startDate = &t
			}
		}
		if e := r.URL.Query().Get("endDate"); e != "" {
			t, err := time.Parse(time.RFC3339, e)
			if err == nil {
				endDate = &t
			}
		}
		if p := r.URL.Query().Get("page"); p != "" {
			if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
				page = parsed
			}
		}
		if l := r.URL.Query().Get("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 30 {
				limit = parsed
			}
		}

		list, err := pvzUC.GetAllWithDetails(startDate, endDate, page, limit)
		if err != nil {
			http.Error(w, "failed to get data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	}
}
