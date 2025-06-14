package service

import (
	"database/sql"
	"user-service/internal/Model"
	"user-service/internal/repository"
	"user-service/internal/request"
)

type UserServiceInterface interface {
	Register(user request.RegisterUser) (*Model.User, error)
}

type UserService struct {
	db             *sql.DB
	UserRepository *repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db:             db,
		UserRepository: repository.NewUserRepository(db),
	}
}

func (u *UserService) Register(user request.RegisterUser) (*Model.User, error) {

	err, data := u.UserRepository.Register(user)
	if err != nil {
		return nil, err
	}

	return data, nil
}
