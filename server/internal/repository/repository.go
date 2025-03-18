package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepository       UserRepository
	AttendanceRepository AttendanceRepository
}

func NewRepositories(DB *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(DB),
		AttendanceRepository: NewAttendanceRepository(DB),
	}
}
