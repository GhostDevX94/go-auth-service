package pkg

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	const defaultCost = 12

	costStr := os.Getenv("BCRYPT_COST")
	intCost := defaultCost
	if costStr != "" {
		parsed, err := strconv.Atoi(costStr)
		if err != nil {
			return "", errors.New("invalid BCRYPT_COST: must be an integer")
		}
		if parsed < bcrypt.MinCost || parsed > bcrypt.MaxCost {
			return "", errors.New("invalid BCRYPT_COST: out of allowed range")
		}
		intCost = parsed
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), intCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
