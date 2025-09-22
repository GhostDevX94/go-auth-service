package service

import "database/sql"

type Services struct {
	UserService UserServiceInterface
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		UserService: NewUserService(db),
	}
}
