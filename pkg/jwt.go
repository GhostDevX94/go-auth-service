package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"user-service/internal/model"
)

func CreateToken(user *model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"iss": "todo-app",
		"aud": user.Email,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
