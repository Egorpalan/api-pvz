package repository

import (
	"errors"
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ReceptionRepository struct {
	db *sqlx.DB
}

func NewReceptionRepository(db *sqlx.DB) *ReceptionRepository {
	return &ReceptionRepository{db: db}
}

func (r *ReceptionRepository) GetLastReception(pvzID string) (*domain.Reception, error) {
	var reception domain.Reception
	err := r.db.Get(&reception, `
        SELECT * FROM receptions 
        WHERE pvz_id = $1 
        ORDER BY date_time DESC 
        LIMIT 1
    `, pvzID)

	if err != nil {
		return nil, nil
	}

	return &reception, nil
}

func (r *ReceptionRepository) Create(reception domain.Reception) (domain.Reception, error) {
	_, err := r.db.Exec(`
        INSERT INTO receptions (id, pvz_id, date_time, status) 
        VALUES ($1, $2, $3, $4)
    `, reception.ID, reception.PVZID, reception.DateTime, reception.Status)

	return reception, err
}

func (r *ReceptionRepository) CloseLastReception(pvzID string) (*domain.Reception, error) {
	var reception domain.Reception
	err := r.db.Get(&reception, `
        SELECT * FROM receptions 
        WHERE pvz_id = $1 AND status = 'in_progress'
        ORDER BY date_time DESC 
        LIMIT 1
    `, pvzID)
	if err != nil {
		return nil, errors.New("no active reception found")
	}

	_, err = r.db.Exec(`
        UPDATE receptions SET status = 'close' 
        WHERE id = $1
    `, reception.ID)
	if err != nil {
		return nil, err
	}

	reception.Status = "close"
	return &reception, nil
}
