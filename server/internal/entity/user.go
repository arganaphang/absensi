package entity

import "time"

const TABLE_USERS = "users"

type UserJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleStaff UserRole = "staff"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Fullname  string    `json:"fullname" db:"fullname"`
	Birthdate time.Time `json:"birthdate" db:"birthdate"`
	Position  string    `json:"position" db:"position"`
	Password  string    `json:"-" db:"password"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	Role      UserRole  `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
