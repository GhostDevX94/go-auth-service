package repository

import "database/sql"

type Repositories struct {
	UserRepository UserRepositoryInterface
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
