package usecase

import (
	"errors"
	"github.com/Egorpalan/api-pvz/internal/metrics"
	"time"

	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
)

type ProductRepo interface {
	GetLastReception(pvzID string) (*domain.Reception, error)
	Create(product domain.Product) (domain.Product, error)
	DeleteLastProductByPVZ(pvzID string) error
}

type ProductUsecase struct {
	repo ProductRepo
}

func NewProductUsecase(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

var allowedTypes = map[string]bool{
	"электроника": true,
	"одежда":      true,
	"обувь":       true,
}

func (u *ProductUsecase) Add(pvzID, productType string) (domain.Product, error) {
	if !allowedTypes[productType] {
		return domain.Product{}, errors.New("invalid product type")
	}

	reception, err := u.repo.GetLastReception(pvzID)
	if err != nil || reception == nil || reception.Status != "in_progress" {
		return domain.Product{}, errors.New("no active reception")
	}

	product := domain.Product{
		ID:          uuid.NewString(),
		ReceptionID: reception.ID,
		DateTime:    time.Now(),
		Type:        productType,
	}

	metrics.ProductsCreated.Inc()

	return u.repo.Create(product)
}

func (u *ProductUsecase) DeleteLast(pvzID string) error {
	return u.repo.DeleteLastProductByPVZ(pvzID)
}
