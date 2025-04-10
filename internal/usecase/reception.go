package usecase

import (
	"errors"
	"github.com/Egorpalan/api-pvz/internal/metrics"
	"time"

	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
)

type ReceptionRepo interface {
	GetLastReception(pvzID string) (*domain.Reception, error)
	Create(reception domain.Reception) (domain.Reception, error)
	CloseLastReception(pvzID string) (*domain.Reception, error)
}

type ReceptionUsecase struct {
	repo ReceptionRepo
}

func NewReceptionUsecase(r ReceptionRepo) *ReceptionUsecase {
	return &ReceptionUsecase{repo: r}
}

func (u *ReceptionUsecase) Create(pvzID string) (domain.Reception, error) {
	last, err := u.repo.GetLastReception(pvzID)
	if err != nil {
		return domain.Reception{}, err
	}

	if last != nil && last.Status == "in_progress" {
		return domain.Reception{}, errors.New("reception is already open")
	}

	reception := domain.Reception{
		ID:       uuid.NewString(),
		PVZID:    pvzID,
		DateTime: time.Now(),
		Status:   "in_progress",
	}

	metrics.ReceptionsCreated.Inc()

	return u.repo.Create(reception)
}

func (u *ReceptionUsecase) CloseLast(pvzID string) (domain.Reception, error) {
	rec, err := u.repo.CloseLastReception(pvzID)
	if err != nil {
		return domain.Reception{}, err
	}
	return *rec, nil
}
