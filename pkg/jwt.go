package pkg

import (
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
