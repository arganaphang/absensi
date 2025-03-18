package repository

import (
	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/internal/middleware"
	"absensi/pkg"
	"context"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type AttendanceRepository interface {
	GetAttendances(ctx context.Context, request dto.GetAttendancesRequest) ([]entity.Attendance, entity.Meta, error)
	GetTodayAttendance(ctx context.Context) ([]entity.Attendance, error)
	CreateAttendance(c context.Context, data dto.CreateAttendanceRequest) (*entity.Attendance, error)
	ApproveAttendance(c context.Context, id string) (*entity.Attendance, error)
}

type attendanceRepository struct {
	DB *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) AttendanceRepository {
	return &attendanceRepository{
		DB: db,
	}
}

func (r attendanceRepository) GetAttendances(ctx context.Context, request dto.GetAttendancesRequest) ([]entity.Attendance, entity.Meta, error) {
	meta := entity.Meta{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   0,
	}
	limit, offset := meta.ToSQL()
	stmt := goqu.From(entity.TABLE_ATTENDANCES).Limit(limit).Offset(offset)
	if request.StartDate != "" { // TODO: Validate Date
		stmt = stmt.Where(goqu.C("created_at").Gte(request.StartDate))
	}
	if request.EndDate != "" { // TODO: Validate Date + 1 Day
		stmt = stmt.Where(goqu.C("created_at").Lt(request.EndDate))
	}
	sql, _, _ := stmt.Order(goqu.C("created_at").Desc()).ToSQL()
	attendances := []entity.Attendance{}
	if err := r.DB.Select(&attendances, sql); err != nil {
		return nil, meta, err
	}
	sqlCount, _, err := stmt.Select(goqu.COUNT("*")).ToSQL()
	if err != nil {
		return nil, meta, err
	}

	var count uint
	if err := r.DB.Get(&count, sqlCount); err != nil {
		return nil, meta, err
	}
	meta.SetTotal(count)
	return attendances, meta, nil
}

func (r attendanceRepository) GetTodayAttendance(ctx context.Context) ([]entity.Attendance, error) {
	today := time.Now().Format(time.DateOnly)
	stmt := goqu.From(entity.TABLE_ATTENDANCES)
	sql, _, _ := stmt.Where(goqu.L("DATE(created_at)").Eq(today)).Order(goqu.C("created_at").Desc()).ToSQL()
	attendances := []entity.Attendance{}
	if err := r.DB.Select(&attendances, sql); err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r attendanceRepository) CreateAttendance(c context.Context, data dto.CreateAttendanceRequest) (*entity.Attendance, error) {
	user, ok := c.Value(middleware.JWTContextUserKey).(*entity.User)
	if !ok {
		return nil, pkg.ErrFailedParseUserAuth
	}
	fmt.Println(user.ID)
	stmt := goqu.Insert(entity.TABLE_ATTENDANCES).Rows(
		goqu.Record{
			"user_id":   user.ID,
			"type":      data.Type,
			"longitude": data.Longitude,
			"latitude":  data.Latitude,
			"note":      data.Note,
		},
	).Returning("*")
	sql, _, _ := stmt.ToSQL()
	attendance := entity.Attendance{}
	if err := r.DB.Get(&attendance, sql); err != nil {
		return nil, err
	}
	return &attendance, nil

}

func (r attendanceRepository) ApproveAttendance(c context.Context, id string) (*entity.Attendance, error) {
	panic("unimplemented")
}
