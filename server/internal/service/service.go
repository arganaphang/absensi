package service

import "absensi/internal/repository"

type Services struct {
	UserService UserService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		UserService: NewUserService(repositories),
	}
}
