package repository

import (
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PVZRepository struct {
	db *sqlx.DB
}

func NewPVZRepository(db *sqlx.DB) *PVZRepository {
	return &PVZRepository{db: db}
}

func (r *PVZRepository) Create(pvz domain.PVZ) (domain.PVZ, error) {
	query := `INSERT INTO pvz (id, city, registration_date) 
              VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, pvz.ID, pvz.City, pvz.RegistrationDate)
	return pvz, err
}
