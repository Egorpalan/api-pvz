package usecase

import (
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockPVZRepo struct {
	mock.Mock
}

func (m *mockPVZRepo) Create(pvz domain.PVZ) (domain.PVZ, error) {
	args := m.Called(pvz)
	return args.Get(0).(domain.PVZ), args.Error(1)
}

func (m *mockPVZRepo) GetAll() ([]domain.PVZ, error) {
	args := m.Called()
	return args.Get(0).([]domain.PVZ), args.Error(1)
}

func TestPVZUsecase_Create(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name        string
		inputCity   string
		mockSetup   func(m *mockPVZRepo)
		expectError bool
	}{
		{
			name:      "valid city - Moscow",
			inputCity: "Москва",
			mockSetup: func(m *mockPVZRepo) {
				m.On("Create", mock.MatchedBy(func(pvz domain.PVZ) bool {
					return pvz.City == "Москва"
				})).Return(domain.PVZ{
					ID:               uuid.New().String(),
					City:             "Москва",
					RegistrationDate: now,
				}, nil)
			},
			expectError: false,
		},
		{
			name:        "invalid city - Тюмень",
			inputCity:   "Тюмень",
			mockSetup:   func(m *mockPVZRepo) {},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockPVZRepo)
			tt.mockSetup(mockRepo)

			uc := NewPVZUsecase(mockRepo)
			result, err := uc.Create(tt.inputCity)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.inputCity, result.City)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
