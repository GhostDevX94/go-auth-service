package dto

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
