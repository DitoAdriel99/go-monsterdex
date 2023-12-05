package tokenizer

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type JWT struct {
	Cfg    config.Cfg
	Issuer string
}

func (j *JWT) GenerateToken(user entity.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(j.Cfg.JWT.Expired) * time.Second)

	claims := &Claims{
		ID:       user.ID,
		Username: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(j.Cfg.JWT.Key))
	if err != nil {
		log.Println("Signed String error : ", err)
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) GetClaimsFromToken(param string) (*Claims, error) {
	tokenString := strings.TrimPrefix(param, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Cfg.JWT.Key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %s", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		log.Println("Token is Invalid")
		return nil, fmt.Errorf("Invalid token")
	}
}
