package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("email", email).Debug("🔍 Executing database query to find user by email")
	
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
		if err == sql.ErrNoRows {
			logrus.WithField("email", email).Debug("👤 User not found in database")
		} else {
			logrus.WithError(err).WithField("email", email).Error("❌ Database error while finding user")
		}
		return nil, err
	}

	logrus.WithField("user_id", user.Id).Debug("✅ User found in database")
	return &user, nil
}

func (r *UserRepository) Register(ctx context.Context, payload dto.RegisterUser) (*model.User, error) {
	logrus.WithField("email", payload.Email).Debug("💾 Executing database query to register new user")

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
		logrus.WithError(err).WithField("email", payload.Email).Error("❌ Database error while registering user")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.Id,
		"email":   user.Email,
	}).Debug("✅ User successfully registered in database")

	return &user, nil
}
