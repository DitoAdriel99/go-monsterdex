package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Email, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}

type LoginResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type Register struct {
	ID        int        `json:"id"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RegisterPayload struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l RegisterPayload) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.FullName, validation.Required),
		validation.Field(&l.Email, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}
