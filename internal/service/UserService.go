package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("email", user.Email).Info("🔐 Hashing user password")

	hasPassword, err := pkg.HashPassword(user.Password)

	if err != nil {
		logrus.WithError(err).Error("❌ Failed to hash password")
		return nil, err
	}

	user.Password = hasPassword

	logrus.WithField("email", user.Email).Info("💾 Saving user to database")
	data, err := u.UserRepository.Register(ctx, user)
	if err != nil {
		logrus.WithError(err).WithField("email", user.Email).Error("❌ Failed to save user to database")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"email": data.Email,
	}).Info("✅ User saved successfully")

	return data, nil
}

func (u *UserService) Login(ctx context.Context, data dto.LoginUser) (string, error) {
	logrus.WithField("email", data.Email).Info("🔍 Looking up user by email")

	user, err := u.UserRepository.GetByEmail(ctx, data.Email)
	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ User not found")
		return "", err
	}

	logrus.WithField("email", data.Email).Info("🔐 Verifying password")
	hasPassword := pkg.CheckPasswordHash(data.Password, user.Password)

	if !hasPassword {
		logrus.WithField("email", data.Email).Warn("⚠️ Invalid password provided")
		return "", errors.New("password is wrong")
	}

	logrus.WithField("email", data.Email).Info("🎫 Generating JWT token")
	token, err := pkg.CreateToken(user)
	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ Failed to generate JWT token")
		return "", err
	}

	logrus.WithField("email", data.Email).Info("✅ Login successful, token generated")
	return token, nil
}
