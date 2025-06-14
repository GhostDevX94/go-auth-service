package Model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
