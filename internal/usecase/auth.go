package usecase

import (
	"errors"

	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/Egorpalan/api-pvz/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	GetByEmail(email string) (*domain.User, error)
	Create(user domain.User) (domain.User, error)
}

type AuthUsecase struct {
	repo UserRepo
}

func NewAuthUsecase(r UserRepo) *AuthUsecase {
	return &AuthUsecase{repo: r}
}

func (a *AuthUsecase) Register(email, password, role string) (domain.User, error) {
	if role != "employee" && role != "moderator" {
		return domain.User{}, errors.New("invalid role")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: string(hash),
		Role:     role,
	}

	return a.repo.Create(user)
}

func (a *AuthUsecase) Login(email, password string) (string, error) {
	user, err := a.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return jwt.GenerateToken(user.Role)
}
