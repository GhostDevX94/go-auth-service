package request

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	IDNumber string `json:"id_number" validate:"required"`
	Address  string `json:"address"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
