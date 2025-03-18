package dto

import (
	"absensi/internal/entity"
)

type GetAttendancesRequest struct {
	Page      uint   `form:"page"`
	PerPage   uint   `form:"per_page"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

type GetAttendancesResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    GetAttendancesResponseData `json:"data"`
}

type GetAttendancesResponseData struct {
	Attendances []entity.Attendance `json:"attendances"`
	Meta        entity.Meta         `json:"meta"`
}

type GetTodayAttendanceRequest struct{}

type GetTodayAttendanceResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Data    GetTodayAttendanceResponseData `json:"data"`
}

type GetTodayAttendanceResponseData struct {
	Attendances []entity.Attendance `json:"attendances"`
}

type CreateAttendanceRequest struct {
	Type      entity.AttendanceType `json:"type"`
	Longitude float64               `json:"longitude"`
	Latitude  float64               `json:"latitude"`
	Note      *string               `json:"note"`
}

type CreateAttendanceResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Data    CreateAttendanceResponseData `json:"data"`
}

type CreateAttendanceResponseData struct {
	Attendance *entity.Attendance `json:"attendance"`
}

type ApproveAttendanceRequest struct{}

type ApproveAttendanceResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Data    ApproveAttendanceResponseData `json:"data"`
}

type ApproveAttendanceResponseData struct {
	Attendance *entity.Attendance `json:"attendance"`
}
