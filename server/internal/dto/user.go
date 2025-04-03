package dto

import (
	"absensi/internal/entity"
	"time"
)

type UserGetProfileResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    UserGetProfileResponseData `json:"data"`
}

type UserGetProfileResponseData struct {
	User *entity.User `json:"user"`
}

type UserUpdateProfileRequest struct {
	Email     string    `json:"email" binding:"required"`
	Fullname  string    `json:"fullname" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Address   string    `json:"address" binding:"required"`
}

type UserUpdateProfileResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Data    UserUpdateProfileResponseData `json:"data"`
}

type UserUpdateProfileResponseData struct {
	User *entity.User `json:"users"`
}

type PasswordChangeRequest struct {
	Token           string `json:"token" binding:"required"`
	CurrentPassword string `json:"current_password" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RePassword      string `json:"re_password" binding:"required"`
}

type PasswordChangeResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    PasswordChangeResponseData `json:"data"`
}

type PasswordChangeResponseData struct {
	User *entity.User `json:"users"`
}

type UserGetAllRequest struct {
	Page    uint   `form:"page"`
	PerPage uint   `form:"per_page"`
	Q       string `form:"q"`
	Role    string `form:"role"`
	Sort    string `form:"sort"`
	SortBy  string `form:"sort_by"`
}

type UserGetAllResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    UserGetAllResponseData `json:"data"`
}

type UserGetAllResponseData struct {
	Users []entity.User `json:"users"`
	Meta  entity.Meta   `json:"meta"`
}

type UserGetByIDResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Data    UserGetByIDResponseData `json:"data"`
}

type UserGetByIDResponseData struct {
	User *entity.User `json:"user"`
}

type UserCreateRequest struct {
	Email     string          `json:"email" binding:"required"`
	Fullname  string          `json:"fullname" binding:"required"`
	Birthdate time.Time       `json:"birthdate" binding:"required"`
	Position  string          `json:"position" binding:"required"`
	Phone     string          `json:"phone" binding:"required"`
	Address   string          `json:"address" binding:"required"`
	Role      entity.UserRole `json:"role" binding:"required"`
}

type UserCreateResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    UserCreateResponseData `json:"data"`
}

type UserCreateResponseData struct {
	User *entity.User `json:"user"`
}

type UserUpdateRequest struct {
	Email     string          `json:"email" binding:"required"`
	Fullname  string          `json:"fullname" binding:"required"`
	Birthdate time.Time       `json:"birthdate" binding:"required"`
	Position  string          `json:"position" binding:"required"`
	Phone     string          `json:"phone" binding:"required"`
	Address   string          `json:"address" binding:"required"`
	Role      entity.UserRole `json:"role" binding:"required"`
}

type UserUpdateResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    UserUpdateResponseData `json:"data"`
}

type UserUpdateResponseData struct {
	User *entity.User `json:"users"`
}

type UserDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	Token *entity.UserJWT `json:"token"`
	User  *entity.User    `json:"user"`
}

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Data    RefreshTokenResponseData `json:"data"`
}

type RefreshTokenResponseData struct {
	Token *entity.UserJWT `json:"token"`
	User  *entity.User    `json:"user"`
}
