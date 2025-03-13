package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepository UserRepository
}

func NewRepositories(DB *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(DB),
	}
}
