package usecase

import (
	"errors"
	"time"

	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
)

var allowedCities = map[string]bool{
	"Москва":          true,
	"Санкт-Петербург": true,
	"Казань":          true,
}

type PVZRepo interface {
	Create(pvz domain.PVZ) (domain.PVZ, error)
}

type PVZUsecase struct {
	repo PVZRepo
}

func NewPVZUsecase(r PVZRepo) *PVZUsecase {
	return &PVZUsecase{repo: r}
}

func (u *PVZUsecase) Create(city string) (domain.PVZ, error) {
	if !allowedCities[city] {
		return domain.PVZ{}, errors.New("unsupported city")
	}

	pvz := domain.PVZ{
		ID:               uuid.NewString(),
		City:             city,
		RegistrationDate: time.Now(),
	}

	return u.repo.Create(pvz)
}
