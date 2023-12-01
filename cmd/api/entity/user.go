package entity

import (
	"time"
)

type User struct {
	ID        int        `json:"id"`
	FullName  string     `json:"fullname"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsActive  bool       `json:"is_active"`
}
