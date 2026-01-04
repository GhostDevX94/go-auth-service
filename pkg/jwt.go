package pkg

import (
	"errors"
	"os"
	"time"
	"user-service/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateToken(user *model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"iss": "user-service",
		"jti": uuid.NewString(),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// CreateRefreshToken generates a UUID-based refresh token
func CreateRefreshToken() string {
	return uuid.NewString()
}

// CreateTokenPair generates both access and refresh tokens
func CreateTokenPair(user *model.User) (accessToken string, refreshToken string, expiresIn int64, err error) {
	accessToken, err = CreateToken(user)
	if err != nil {
		return "", "", 0, err
	}

	refreshToken = CreateRefreshToken()
	expiresIn = 3600 // 1 hour in seconds

	return accessToken, refreshToken, expiresIn, nil
}

// ValidateToken validates and parses a JWT token
func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}
