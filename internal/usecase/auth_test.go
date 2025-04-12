package usecase

import (
	"errors"
	"testing"

	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) Create(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockUserRepo) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if u := args.Get(0); u != nil {
		return u.(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAuthUsecase_Register(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		role        string
		mockSetup   func(m *mockUserRepo)
		expected    domain.User
		expectError bool
	}{
		{
			name:     "successful registration with client role",
			email:    "test@example.com",
			password: "securepass",
			role:     "client",
			mockSetup: func(m *mockUserRepo) {
				expectedUser := domain.User{
					ID:       "mocked-uuid",
					Email:    "test@example.com",
					Password: "hashed-password",
					Role:     "client",
				}

				m.On("Create", mock.MatchedBy(func(u domain.User) bool {
					err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("securepass"))
					return u.Email == "test@example.com" &&
						u.Role == "client" &&
						err == nil &&
						u.ID != ""
				})).Return(expectedUser, nil)
			},
			expected: domain.User{
				ID:       "mocked-uuid",
				Email:    "test@example.com",
				Password: "hashed-password",
				Role:     "client",
			},
			expectError: false,
		},
		{
			name:     "successful registration with moderator role",
			email:    "mod@example.com",
			password: "modpass",
			role:     "moderator",
			mockSetup: func(m *mockUserRepo) {
				expectedUser := domain.User{
					ID:       "mod-uuid",
					Email:    "mod@example.com",
					Password: "hashed-mod-pass",
					Role:     "moderator",
				}

				m.On("Create", mock.MatchedBy(func(u domain.User) bool {
					return u.Email == "mod@example.com" && u.Role == "moderator"
				})).Return(expectedUser, nil)
			},
			expected: domain.User{
				ID:       "mod-uuid",
				Email:    "mod@example.com",
				Password: "hashed-mod-pass",
				Role:     "moderator",
			},
			expectError: false,
		},
		{
			name:     "invalid role",
			email:    "test@example.com",
			password: "securepass",
			role:     "admin",
			mockSetup: func(m *mockUserRepo) {
			},
			expected:    domain.User{},
			expectError: true,
		},
		{
			name:     "database error",
			email:    "test@example.com",
			password: "securepass",
			role:     "client",
			mockSetup: func(m *mockUserRepo) {
				m.On("Create", mock.MatchedBy(func(u domain.User) bool {
					return u.Email == "test@example.com" && u.Role == "client"
				})).Return(domain.User{}, errors.New("database error"))
			},
			expected:    domain.User{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockUserRepo)
			tt.mockSetup(mockRepo)

			uc := NewAuthUsecase(mockRepo)
			user, err := uc.Register(tt.email, tt.password, tt.role)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, domain.User{}, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestAuthUsecase_Login(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	tests := []struct {
		name        string
		email       string
		password    string
		mockSetup   func(m *mockUserRepo)
		expectToken bool
		expectError bool
	}{
		{
			name:     "successful login",
			email:    "existing@example.com",
			password: "correctpassword",
			mockSetup: func(m *mockUserRepo) {
				m.On("GetByEmail", "existing@example.com").Return(&domain.User{
					ID:       "user-uuid",
					Email:    "existing@example.com",
					Password: string(hashedPassword),
					Role:     "client",
				}, nil)
			},
			expectToken: true,
			expectError: false,
		},
		{
			name:     "user not found",
			email:    "nonexistent@example.com",
			password: "anypassword",
			mockSetup: func(m *mockUserRepo) {
				m.On("GetByEmail", "nonexistent@example.com").Return(nil, errors.New("user not found"))
			},
			expectToken: false,
			expectError: true,
		},
		{
			name:     "incorrect password",
			email:    "existing@example.com",
			password: "wrongpassword",
			mockSetup: func(m *mockUserRepo) {
				m.On("GetByEmail", "existing@example.com").Return(&domain.User{
					ID:       "user-uuid",
					Email:    "existing@example.com",
					Password: string(hashedPassword),
					Role:     "client",
				}, nil)
			},
			expectToken: false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockUserRepo)
			tt.mockSetup(mockRepo)

			uc := NewAuthUsecase(mockRepo)
			token, err := uc.Login(tt.email, tt.password)

			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
