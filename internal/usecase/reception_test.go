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

func TestReceptionUsecase_CloseLast(t *testing.T) {
	pvzID := uuid.NewString()
	activeReception := &domain.Reception{
		ID:     uuid.NewString(),
		PVZID:  pvzID,
		Status: "in_progress",
	}

	closedReception := &domain.Reception{
		ID:     activeReception.ID,
		PVZID:  pvzID,
		Status: "close",
	}

	tests := []struct {
		name        string
		mockSetup   func(m *mockReceptionRepo)
		expectError bool
	}{
		{
			name: "successful close of active reception",
			mockSetup: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(activeReception, nil)
				m.On("CloseLastReception", pvzID).Return(closedReception, nil)
			},
			expectError: false,
		},
		{
			name: "no reception found",
			mockSetup: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(nil, nil)
			},
			expectError: true,
		},
		{
			name: "reception already closed",
			mockSetup: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(&domain.Reception{
					ID:     uuid.NewString(),
					PVZID:  pvzID,
					Status: "close",
				}, nil)
			},
			expectError: true,
		},
		{
			name: "repository GetLastReception error",
			mockSetup: func(m *mockReceptionRepo) {
				m.On("GetLastReception", pvzID).Return(nil, errors.New("db error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockReceptionRepo)
			tt.mockSetup(mockRepo)

			uc := NewReceptionUsecase(mockRepo)
			_, err := uc.CloseLast(pvzID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
