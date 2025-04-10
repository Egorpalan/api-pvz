package repository

import (
	"errors"
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetLastReception(pvzID string) (*domain.Reception, error) {
	var rec domain.Reception
	err := r.db.Get(&rec, `
        SELECT * FROM receptions 
        WHERE pvz_id = $1 
        ORDER BY date_time DESC 
        LIMIT 1
    `, pvzID)

	if err != nil {
		return nil, nil
	}
	return &rec, nil
}

func (r *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	_, err := r.db.Exec(`
        INSERT INTO products (id, reception_id, date_time, type)
        VALUES ($1, $2, $3, $4)
    `, product.ID, product.ReceptionID, product.DateTime, product.Type)
	return product, err
}

func (r *ProductRepository) DeleteLastProductByPVZ(pvzID string) error {
	var rec domain.Reception
	err := r.db.Get(&rec, `
        SELECT * FROM receptions 
        WHERE pvz_id = $1 AND status = 'in_progress'
        ORDER BY date_time DESC 
        LIMIT 1
    `, pvzID)
	if err != nil {
		return errors.New("no active reception")
	}

	var product domain.Product
	err = r.db.Get(&product, `
        SELECT * FROM products 
        WHERE reception_id = $1
        ORDER BY date_time DESC 
        LIMIT 1
    `, rec.ID)
	if err != nil {
		return errors.New("no products to delete")
	}

	_, err = r.db.Exec(`DELETE FROM products WHERE id = $1`, product.ID)
	return err
}
