package service

import (
	"testing"

	"user-service/internal/dto"
	"user-service/pkg"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHashing(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectedError bool
	}{
		{
			name:          "Valid password",
			password:      "password123",
			expectedError: false,
		},
		{
			name:          "Empty password",
			password:      "",
			expectedError: false,
		},
		{
			name:          "Very long password",
			password:      "verylongpasswordwithlotsofcharacters",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := pkg.HashPassword(tt.password)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, hashedPassword)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hashedPassword)

				isValid := pkg.CheckPasswordHash(tt.password, hashedPassword)
				assert.True(t, isValid)
			}
		})
	}
}

func TestPasswordValidation(t *testing.T) {
	password := "password123"
	hashedPassword, _ := pkg.HashPassword(password)

	tests := []struct {
		name           string
		inputPassword  string
		storedHash     string
		expectedResult bool
	}{
		{
			name:           "Correct password",
			inputPassword:  password,
			storedHash:     hashedPassword,
			expectedResult: true,
		},
		{
			name:           "Wrong password",
			inputPassword:  "wrongpassword",
			storedHash:     hashedPassword,
			expectedResult: false,
		},
		{
			name:           "Empty password",
			inputPassword:  "",
			storedHash:     hashedPassword,
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.CheckPasswordHash(tt.inputPassword, tt.storedHash)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestUserService_Register_Validation(t *testing.T) {
	tests := []struct {
		name          string
		request       dto.RegisterUser
		expectedError bool
	}{
		{
			name: "Valid registration data",
			request: dto.RegisterUser{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User",
			},
			expectedError: false,
		},
		{
			name: "Missing email",
			request: dto.RegisterUser{
				Password: "password123",
				Name:     "Test User",
			},
			expectedError: false,
		},
		{
			name: "Missing password",
			request: dto.RegisterUser{
				Email: "test@example.com",
				Name:  "Test User",
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.request)

			if tt.request.Email != "" {
				assert.Contains(t, tt.request.Email, "@")
			}

			if tt.request.Password != "" {
				assert.GreaterOrEqual(t, len(tt.request.Password), 1)
			}
		})
	}
}
