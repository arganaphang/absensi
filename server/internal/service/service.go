package service

import "absensi/internal/repository"

type Services struct {
	UserService       UserService
	AttendanceService AttendanceService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		UserService:       NewUserService(repositories),
		AttendanceService: NewAttendanceService(repositories),
	}
}
