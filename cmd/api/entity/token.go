package entity

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKEY = []byte(os.Getenv("KEY"))

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
