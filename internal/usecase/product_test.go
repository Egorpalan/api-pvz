package usecase

import (
	"errors"
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockProductRepo struct {
	mock.Mock
}

func (m *mockProductRepo) GetLastReception(pvzID string) (*domain.Reception, error) {
	args := m.Called(pvzID)
	if res := args.Get(0); res != nil {
		return res.(*domain.Reception), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockProductRepo) Create(product domain.Product) (domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *mockProductRepo) DeleteLastProductByPVZ(pvzID string) error {
	args := m.Called(pvzID)
	return args.Error(0)
}

func TestProductUsecase_Add(t *testing.T) {
	now := time.Now()
	pvzID := uuid.NewString()
	receptionID := uuid.NewString()

	tests := []struct {
		name        string
		productType string
		mockSetup   func(m *mockProductRepo)
		expectError bool
	}{
		{
			name:        "good add electric",
			productType: "электроника",
			mockSetup: func(m *mockProductRepo) {
				m.On("GetLastReception", pvzID).Return(&domain.Reception{
					ID:     receptionID,
					PVZID:  pvzID,
					Status: "in_progress",
				}, nil)

				m.On("Create", mock.AnythingOfType("domain.Product")).Return(domain.Product{
					ID:          uuid.NewString(),
					Type:        "электроника",
					DateTime:    now,
					ReceptionID: receptionID,
				}, nil)
			},
			expectError: false,
		},
		{
			name:        "wrong type product",
			productType: "мебель",
			mockSetup:   func(m *mockProductRepo) {},
			expectError: true,
		},
		{
			name:        "no active reception",
			productType: "одежда",
			mockSetup: func(m *mockProductRepo) {
				m.On("GetLastReception", pvzID).Return(nil, nil)
			},
			expectError: true,
		},
		{
			name:        "error reception",
			productType: "обувь",
			mockSetup: func(m *mockProductRepo) {
				m.On("GetLastReception", pvzID).Return(nil, errors.New("db error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockProductRepo)
			tt.mockSetup(mockRepo)

			uc := NewProductUsecase(mockRepo)
			_, err := uc.Add(pvzID, tt.productType)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestProductUsecase_DeleteLast(t *testing.T) {
	pvzID := uuid.NewString()

	tests := []struct {
		name        string
		mockSetup   func(m *mockProductRepo)
		expectError bool
	}{
		{
			name: "successful deletion",
			mockSetup: func(m *mockProductRepo) {
				m.On("DeleteLastProductByPVZ", pvzID).Return(nil)
			},
			expectError: false,
		},
		{
			name: "deletion error from repo",
			mockSetup: func(m *mockProductRepo) {
				m.On("DeleteLastProductByPVZ", pvzID).Return(errors.New("db error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockProductRepo)
			tt.mockSetup(mockRepo)

			uc := NewProductUsecase(mockRepo)
			err := uc.DeleteLast(pvzID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
