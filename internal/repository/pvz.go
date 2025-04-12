package repository

import (
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/Egorpalan/api-pvz/internal/dto"
	"github.com/jmoiron/sqlx"
	"time"
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

func (r *PVZRepository) GetWithReceptionsAndProducts(start, end *time.Time, page, limit int) ([]dto.PVZDTO, error) {
	offset := (page - 1) * limit

	query := `
        SELECT * FROM pvz
        WHERE ($1::timestamptz IS NULL OR registration_date >= $1)
          AND ($2::timestamptz IS NULL OR registration_date <= $2)
        ORDER BY registration_date DESC
        LIMIT $3 OFFSET $4
    `
	var pvzList []domain.PVZ
	if err := r.db.Select(&pvzList, query, start, end, limit, offset); err != nil {
		return nil, err
	}

	var result []dto.PVZDTO

	for _, pvz := range pvzList {
		dtoPVZ := dto.PVZDTO{
			ID:               pvz.ID,
			City:             pvz.City,
			RegistrationDate: pvz.RegistrationDate,
		}

		var receptions []domain.Reception
		err := r.db.Select(&receptions, `SELECT * FROM receptions WHERE pvz_id = $1`, pvz.ID)
		if err != nil {
			return nil, err
		}

		for _, rec := range receptions {
			recDTO := dto.ReceptionDTO{
				ID:       rec.ID,
				DateTime: rec.DateTime,
				Status:   rec.Status,
			}

			var products []domain.Product
			err = r.db.Select(&products, `SELECT * FROM products WHERE reception_id = $1`, rec.ID)
			if err != nil {
				return nil, err
			}

			for _, p := range products {
				recDTO.Products = append(recDTO.Products, dto.ProductDTO{
					ID:       p.ID,
					Type:     p.Type,
					DateTime: p.DateTime,
				})
			}

			dtoPVZ.Receptions = append(dtoPVZ.Receptions, recDTO)
		}

		result = append(result, dtoPVZ)
	}

	return result, nil
}

func (r *PVZRepository) GetAll() ([]domain.PVZ, error) {
	var pvzs []domain.PVZ
	err := r.db.Select(&pvzs, "SELECT * FROM pvz")
	return pvzs, err
}
