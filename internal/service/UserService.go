package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"user-service/internal/dto"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/pkg"

	"github.com/sirupsen/logrus"
)

type UserServiceInterface interface {
	Register(context.Context, dto.RegisterUser) (*model.User, error)
	Login(context.Context, dto.LoginUser) (*dto.ResponseToken, error)
	UpdateToken(context.Context, string) (*dto.ResponseToken, error)
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(db),
	}
}

func (u *UserService) Register(ctx context.Context, user dto.RegisterUser) (*model.User, error) {

	hasPassword, err := pkg.HashPassword(user.Password)

	if err != nil {
		logrus.WithError(err).Error("Failed to hash password")
		return nil, err
	}

	user.Password = hasPassword

	data, err := u.UserRepository.Register(ctx, user)
	if err != nil {
		logrus.WithError(err).WithField("email", user.Email).Error("Failed to save user to database")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"email": data.Email,
	}).Info("User saved successfully")

	return data, nil
}

func (u *UserService) Login(ctx context.Context, data dto.LoginUser) (*dto.ResponseToken, error) {
	logrus.WithField("email", data.Email).Info("🔍 Looking up user by email")

	user, err := u.UserRepository.GetByEmail(ctx, data.Email)
	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("User not found")
		return nil, err
	}

	logrus.WithField("email", data.Email).Info("Verifying password")
	hasPassword := pkg.CheckPasswordHash(data.Password, user.Password)

	if !hasPassword {
		logrus.WithField("email", data.Email).Warn("Invalid password provided")
		return nil, errors.New("password is wrong")
	}

	logrus.WithField("email", data.Email).Info("Generating token pair")
	accessToken, refreshToken, expiresIn, err := pkg.CreateTokenPair(user)
	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("Failed to generate token pair")
		return nil, err
	}

	// Save refresh token to database
	refreshExpiry := time.Now().Add(7 * 24 * time.Hour) // 7 days
	err = u.UserRepository.SaveRefreshToken(ctx, user.Id, refreshToken, refreshExpiry)
	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("Failed to save refresh token")
		return nil, err
	}

	logrus.WithField("email", data.Email).Info("Login successful, token pair generated")
	return &dto.ResponseToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// UpdateToken refreshes the access token using a valid refresh token
func (u *UserService) UpdateToken(ctx context.Context, refreshTokenStr string) (*dto.ResponseToken, error) {
	logrus.Info("Processing token refresh request")

	// Get refresh token from database
	refreshToken, err := u.UserRepository.GetRefreshToken(ctx, refreshTokenStr)
	if err != nil {
		logrus.WithError(err).Error("Refresh token not found")
		return nil, errors.New("invalid refresh token")
	}

	// Check if token is revoked
	if refreshToken.IsRevoked {
		logrus.Warn("Attempted to use revoked refresh token")
		return nil, errors.New("refresh token has been revoked")
	}

	// Check if token is expired
	if time.Now().After(refreshToken.ExpiresAt) {
		logrus.Warn("Attempted to use expired refresh token")
		return nil, errors.New("refresh token has expired")
	}

	// Get user by ID
	user, err := u.UserRepository.GetByEmail(ctx, "") // We need to get by ID instead
	if err != nil {
		logrus.WithError(err).WithField("user_id", refreshToken.UserId).Error("User not found")
		return nil, errors.New("user not found")
	}

	// Generate new token pair
	accessToken, newRefreshToken, expiresIn, err := pkg.CreateTokenPair(user)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate new token pair")
		return nil, err
	}

	// Revoke old refresh token
	err = u.UserRepository.RevokeRefreshToken(ctx, refreshTokenStr)
	if err != nil {
		logrus.WithError(err).Warn("Failed to revoke old refresh token")
		// Continue anyway, as we've already generated new tokens
	}

	// Save new refresh token
	refreshExpiry := time.Now().Add(7 * 24 * time.Hour)
	err = u.UserRepository.SaveRefreshToken(ctx, user.Id, newRefreshToken, refreshExpiry)
	if err != nil {
		logrus.WithError(err).Error("Failed to save new refresh token")
		return nil, err
	}

	logrus.Info("Token refresh successful")
	return &dto.ResponseToken{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}
