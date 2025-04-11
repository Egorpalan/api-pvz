package usecase

import (
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockReceptionRepo struct {
	mock.Mock
}

func (m *mockReceptionRepo) GetLastReception(pvzID string) (*domain.Reception, error) {
	args := m.Called(pvzID)
	if r := args.Get(0); r != nil {
		return r.(*domain.Reception), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockReceptionRepo) Create(reception domain.Reception) (domain.Reception, error) {
	args := m.Called(reception)
	return args.Get(0).(domain.Reception), args.Error(1)
}

func (m *mockReceptionRepo) CloseLastReception(pvzID string) (*domain.Reception, error) {
	args := m.Called(pvzID)
	if r := args.Get(0); r != nil {
		return r.(*domain.Reception), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestReceptionUsecase_Create(t *testing.T) {
	now := time.Now()
	pvzID := uuid.NewString()

	tests := []struct {
		name        string
		setupMock   func(m *mockReceptionRepo)
		expectError bool
	}{
		{
			name: "reception already exists",
			setupMock: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(&domain.Reception{
					ID:     uuid.NewString(),
					PVZID:  pvzID,
					Status: "in_progress",
				}, nil)
			},
			expectError: true,
		},
		{
			name: "reception close, make new",
			setupMock: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(&domain.Reception{
					ID:     uuid.NewString(),
					PVZID:  pvzID,
					Status: "close",
				}, nil)

				m.On("Create", mock.AnythingOfType("domain.Reception")).Return(domain.Reception{
					ID:       uuid.NewString(),
					PVZID:    pvzID,
					DateTime: now,
					Status:   "in_progress",
				}, nil)
			},
			expectError: false,
		},
		{
			name: "no reception, make new",
			setupMock: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(nil, nil)

				m.On("Create", mock.AnythingOfType("domain.Reception")).Return(domain.Reception{
					ID:       uuid.NewString(),
					PVZID:    pvzID,
					DateTime: now,
					Status:   "in_progress",
				}, nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockReceptionRepo)
			tt.setupMock(mockRepo)

			uc := NewReceptionUsecase(mockRepo)
			_, err := uc.Create(pvzID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
