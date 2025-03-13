package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/pkg"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UserRepository interface {
	GetAll(ctx context.Context, request dto.UserGetAllRequest) ([]entity.User, entity.Meta, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error)
	Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r userRepository) GetAll(ctx context.Context, request dto.UserGetAllRequest) ([]entity.User, entity.Meta, error) {
	meta := entity.Meta{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   0,
	}
	limit, offset := meta.ToSQL()
	stmt := goqu.From(entity.TABLE_USERS).Limit(limit).Offset(offset)
	if request.Q != "" {
		q := fmt.Sprintf("%%%s%%", request.Q)
		stmt = stmt.Where(goqu.Or(
			goqu.C("fullname").Like(q),
			goqu.C("email").Like(q),
			goqu.C("phone").Like(q),
		))
	}
	if request.Role != "" {
		stmt = stmt.Where(goqu.C("role").Eq(request.Role))
	}
	if request.Sort != "" {
		if strings.ToLower(request.SortBy) == "asc" {
			stmt = stmt.Order(goqu.C(request.Sort).Asc())
		}
		if strings.ToLower(request.SortBy) == "desc" {
			stmt = stmt.Order(goqu.C(request.Sort).Desc())
		}
	}
	sql, _, _ := stmt.ToSQL()
	users := []entity.User{}
	if err := r.DB.Select(&users, sql); err != nil {
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
	return users, meta, nil
}

func (r userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	sql, _, _ := goqu.From(entity.TABLE_USERS).Where(goqu.C("id").Eq(id)).ToSQL()
	user := entity.User{}
	if err := r.DB.Get(&user, sql); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	sql, _, _ := goqu.From(entity.TABLE_USERS).Where(goqu.C("email").Eq(email)).ToSQL()
	user := entity.User{}
	if err := r.DB.Get(&user, sql); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error) {
	password, err := pkg.HashCreate(data.Birthdate.Format(time.DateOnly)) // "2006-01-02"
	if err != nil {
		return nil, err
	}
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	sql, _, _ := goqu.Insert(entity.TABLE_USERS).
		Cols("id", "fullname", "email", "password", "birthdate", "position", "phone", "address", "role").
		Vals(goqu.Vals{id, data.Fullname, data.Email, password, data.Birthdate, data.Position, data.Phone, data.Address, data.Role}).
		Returning("*").
		ToSQL()
	rows, err := r.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	var user entity.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r userRepository) Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error) {
	record := goqu.Record{
		"fullname":   data.Fullname,
		"email":      data.Email,
		"birthdate":  data.Birthdate,
		"position":   data.Position,
		"phone":      data.Phone,
		"address":    data.Address,
		"role":       data.Role,
		"updated_at": time.Now(),
	}
	sql, _, err := goqu.Update(entity.TABLE_USERS).
		Set(record).
		Where(goqu.C("id").Eq(id)).
		Returning("*").
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	var user entity.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	sql, _, _ := goqu.Delete(entity.TABLE_USERS).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
