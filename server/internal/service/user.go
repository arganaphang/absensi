package service

import (
	"context"

	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/internal/repository"
	"absensi/pkg"
)

type UserService interface {
	GetAll(ctx context.Context, request dto.UserGetAllRequest) ([]entity.User, entity.Meta, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error)
	Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error)
	Delete(ctx context.Context, id string) error

	Login(ctx context.Context, data dto.LoginRequest) (*entity.User, *entity.UserJWT, error)
	RefreshToken(ctx context.Context, data dto.RefreshTokenRequest) (*entity.User, *entity.UserJWT, error)
	UpdatePassword(ctx context.Context, id string, data dto.PasswordChangeRequest) (*entity.User, error)
}

type userService struct {
	Repositories *repository.Repositories
}

func NewUserService(repositories *repository.Repositories) UserService {
	return &userService{
		Repositories: repositories,
	}
}

func (s userService) GetAll(ctx context.Context, request dto.UserGetAllRequest) ([]entity.User, entity.Meta, error) {
	return s.Repositories.UserRepository.GetAll(ctx, request)
}

func (s userService) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.Repositories.UserRepository.GetByID(ctx, id)
}

func (s userService) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.Repositories.UserRepository.GetByEmail(ctx, email)
}

func (s userService) Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error) {
	return s.Repositories.UserRepository.Create(ctx, data)
}

func (s userService) Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error) {
	return s.Repositories.UserRepository.Update(ctx, id, data)
}

func (s userService) Delete(ctx context.Context, id string) error {
	return s.Repositories.UserRepository.Delete(ctx, id)
}

func (s userService) Login(ctx context.Context, data dto.LoginRequest) (*entity.User, *entity.UserJWT, error) {
	user, err := s.Repositories.UserRepository.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, err
	}
	if !pkg.HashCompare(data.Password, user.Password) {
		return nil, nil, pkg.ErrUnauthorized
	}
	token, err := pkg.JWTEncode(*user)
	if err != nil {
		return nil, nil, err
	}
	return user, token, nil
}

func (s userService) RefreshToken(ctx context.Context, data dto.RefreshTokenRequest) (*entity.User, *entity.UserJWT, error) {
	user, token, err := pkg.JWTRefresh(data.AccessToken, data.RefreshToken)
	if err != nil {
		return nil, nil, err
	}
	return user, token, nil
}

func (s userService) UpdatePassword(ctx context.Context, id string, data dto.PasswordChangeRequest) (*entity.User, error) {
	user, err := s.Repositories.UserRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !pkg.HashCompare(data.CurrentPassword, user.Password) {
		return nil, pkg.ErrUnauthorized
	}
	return s.Repositories.UserRepository.UpdatePassword(ctx, id, data)
}
