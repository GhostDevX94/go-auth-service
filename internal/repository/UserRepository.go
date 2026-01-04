package repository

import (
	"context"
	"database/sql"
	"time"
	"user-service/internal/dto"
	"user-service/internal/model"

	"github.com/sirupsen/logrus"
)

type UserRepositoryInterface interface {
	Register(context.Context, dto.RegisterUser) (*model.User, error)
	GetByEmail(context.Context, string) (*model.User, error)
	GetById(context.Context, int) (*model.User, error)
	SaveRefreshToken(context.Context, int, string, time.Time) error
	GetRefreshToken(context.Context, string) (*model.RefreshToken, error)
	RevokeRefreshToken(context.Context, string) error
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
	// add per-query timeout
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	err := r.db.QueryRowContext(qctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logrus.WithField("email", email).Debug("User not found in database")
		} else {
			logrus.WithError(err).WithField("email", email).Error("Database error while finding user")
		}
		return nil, err
	}

	return &user, nil
}

// GetById retrieves a user by their ID
func (r *UserRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	logrus.WithField("user_id", id).Debug("🔍 Executing database query to find user by ID")

	var user model.User
	query := "SELECT id, name, surname, email, phone, password FROM users WHERE id = $1"
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(qctx, query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logrus.WithField("user_id", id).Debug("User not found in database")
		} else {
			logrus.WithError(err).WithField("user_id", id).Error("Database error while finding user")
		}
		return nil, err
	}

	logrus.WithField("user_id", user.Id).Debug("User found in database")
	return &user, nil
}

func (r *UserRepository) Register(ctx context.Context, payload dto.RegisterUser) (*model.User, error) {
	logrus.WithField("email", payload.Email).Debug("Executing database query to register new user")

	var user model.User

	query := "INSERT INTO users (name,surname ,email, phone, password) VALUES ($1, $2, $3, $4,$5) RETURNING id,name,surname ,email, phone"
	// add per-query timeout
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	err := r.db.QueryRowContext(qctx, query, payload.Name, payload.Surname, payload.Email, payload.Phone, payload.Password).Scan(
		&user.Id,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Phone,
	)
	if err != nil {
		logrus.WithError(err).WithField("email", payload.Email).Error("Database error while registering user")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.Id,
		"email":   user.Email,
	}).Debug("User successfully registered in database")

	return &user, nil
}

// SaveRefreshToken stores a new refresh token in the database
func (r *UserRepository) SaveRefreshToken(ctx context.Context, userId int, token string, expiresAt time.Time) error {
	logrus.WithFields(logrus.Fields{
		"user_id": userId,
	}).Debug("Saving refresh token to database")

	query := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(qctx, query, userId, token, expiresAt)
	if err != nil {
		logrus.WithError(err).WithField("user_id", userId).Error("Failed to save refresh token")
		return err
	}

	logrus.WithField("user_id", userId).Debug("Refresh token saved successfully")
	return nil
}

// GetRefreshToken retrieves a refresh token by token string
func (r *UserRepository) GetRefreshToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	logrus.Debug("Fetching refresh token from database")

	var refreshToken model.RefreshToken
	query := "SELECT id, user_id, token, expires_at, created_at, is_revoked FROM refresh_tokens WHERE token = $1"
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(qctx, query, token).Scan(
		&refreshToken.Id,
		&refreshToken.UserId,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
		&refreshToken.IsRevoked,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Debug("Refresh token not found in database")
		} else {
			logrus.WithError(err).Error("Database error while fetching refresh token")
		}
		return nil, err
	}

	logrus.WithField("token_id", refreshToken.Id).Debug("Refresh token found")
	return &refreshToken, nil
}

// RevokeRefreshToken marks a refresh token as revoked
func (r *UserRepository) RevokeRefreshToken(ctx context.Context, token string) error {
	logrus.Debug("Revoking refresh token")

	query := "UPDATE refresh_tokens SET is_revoked = TRUE WHERE token = $1"
	qctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(qctx, query, token)
	if err != nil {
		logrus.WithError(err).Error("Failed to revoke refresh token")
		return err
	}

	logrus.Debug("Refresh token revoked successfully")
	return nil
}
