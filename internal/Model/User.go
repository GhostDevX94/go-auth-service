package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
