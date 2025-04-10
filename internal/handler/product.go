package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/Egorpalan/api-pvz/internal/usecase"
)

type createProductRequest struct {
	Type  string `json:"type"`
	PVZID string `json:"pvzId"`
}

func CreateProduct(uc *usecase.ProductUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		product, err := uc.Add(req.PVZID, req.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func DeleteLastProduct(uc *usecase.ProductUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pvzID := chi.URLParam(r, "pvzId")
		if pvzID == "" {
			http.Error(w, "missing pvzId", http.StatusBadRequest)
			return
		}

		err := uc.DeleteLast(pvzID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("last product deleted"))
	}
}
