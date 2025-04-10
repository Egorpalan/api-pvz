package repository

import (
	"github.com/Egorpalan/api-pvz/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user domain.User) (domain.User, error) {
	_, err := r.db.Exec(`
        INSERT INTO users (id, email, password, role)
        VALUES ($1, $2, $3, $4)
    `, user.ID, user.Email, user.Password, user.Role)
	return user, err
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var u domain.User
	err := r.db.Get(&u, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
