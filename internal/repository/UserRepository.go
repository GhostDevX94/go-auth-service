package repository

import (
	"context"
	"database/sql"
	"user-service/internal/dto"
	"user-service/internal/model"
)

type UserRepositoryInterface interface {
	Register(context.Context, dto.RegisterUser) (*model.User, error)
	GetByEmail(context.Context, string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := "SELECT id,name,surname ,email, phone, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Register(ctx context.Context, payload dto.RegisterUser) (*model.User, error) {

	var user model.User

	query := "INSERT INTO users (name,surname ,email, phone, password) VALUES ($1, $2, $3, $4,$5) RETURNING id,name,surname ,email, phone"
	err := r.db.QueryRowContext(ctx, query, payload.Name, payload.Surname, payload.Email, payload.Phone, payload.Password).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
