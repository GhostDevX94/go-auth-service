package dto

type RegisterUser struct {
	Name     string `json:"name" validate:"required" example:"John"`
	Surname  string `json:"surname" validate:"required" example:"Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Phone    string `json:"phone" validate:"required" example:"+1234567890"`
	Password string `json:"password" validate:"required" example:"SecurePass123!"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example:"SecurePass123!"`
}

type ResponseToken struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresIn    int64  `json:"expires_in" example:"3600"` // in seconds
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
