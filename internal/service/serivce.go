package service

import "database/sql"

type Services struct {
	UserService UserService
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		UserService: *NewUserService(db),
	}
}
