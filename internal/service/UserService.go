package service

import (
	"context"
	"database/sql"
	"errors"
	"user-service/internal/dto"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/pkg"
)

type UserServiceInterface interface {
	Register(context.Context, dto.RegisterUser) (*model.User, error)
}

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(db),
	}
}

func (u *UserService) Register(ctx context.Context, user dto.RegisterUser) (*model.User, error) {

	hasPassword, err := pkg.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hasPassword

	data, err := u.UserRepository.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Login(ctx context.Context, data dto.LoginUser) (*model.User, error) {

	user, err := u.UserRepository.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	hasPassword := pkg.CheckPasswordHash(data.Password, user.Password)

	if !hasPassword {
		return nil, errors.New("password is wrong")
	}

	return user, nil
}
