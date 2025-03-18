package entity

import "time"

const TABLE_ATTENDANCES = "attendances"

type AttendanceType string

const (
	AttendanceTypeIN  AttendanceType = "in"
	AttendanceTypeOUT AttendanceType = "out"
)

type Attendance struct {
	ID        string         `json:"id" db:"id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	Type      AttendanceType `json:"type" db:"type"`
	UserID    string         `json:"user_id" db:"user_id"`
	Longitude float64        `json:"longitude" db:"longitude"`
	Latitude  float64        `json:"latitude" db:"latitude"`
	Note      *string        `json:"note" db:"note"`
}
