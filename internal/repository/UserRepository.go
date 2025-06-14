package repository

import (
	"database/sql"
	"errors"
	"user-service/internal/Model"
	"user-service/internal/request"
)

type UserRepositoryInterface interface {
	Register(user request.RegisterUser) (*Model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Register(user request.RegisterUser) (*Model.User, error) {

	return nil, errors.New("user not created")

}
