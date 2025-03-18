package service

import (
	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/internal/repository"
	"context"
)

type AttendanceService interface {
	GetAttendances(ctx context.Context, request dto.GetAttendancesRequest) ([]entity.Attendance, entity.Meta, error)
	GetTodayAttendance(ctx context.Context) ([]entity.Attendance, error)
	CreateAttendance(c context.Context, data dto.CreateAttendanceRequest) (*entity.Attendance, error)
	ApproveAttendance(c context.Context, id string) (*entity.Attendance, error)
}

type attendanceService struct {
	Repositories *repository.Repositories
}

func NewAttendanceService(repositories *repository.Repositories) AttendanceService {
	return &attendanceService{
		Repositories: repositories,
	}
}

func (s attendanceService) GetAttendances(ctx context.Context, request dto.GetAttendancesRequest) ([]entity.Attendance, entity.Meta, error) {
	return s.Repositories.AttendanceRepository.GetAttendances(ctx, request)
}

func (s attendanceService) GetTodayAttendance(ctx context.Context) ([]entity.Attendance, error) {
	return s.Repositories.AttendanceRepository.GetTodayAttendance(ctx)
}

func (s attendanceService) CreateAttendance(ctx context.Context, data dto.CreateAttendanceRequest) (*entity.Attendance, error) {
	return s.Repositories.AttendanceRepository.CreateAttendance(ctx, data)
}

func (s attendanceService) ApproveAttendance(ctx context.Context, id string) (*entity.Attendance, error) {
	return s.Repositories.AttendanceRepository.ApproveAttendance(ctx, id)

}
