package auth

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/pkg/hashing"

	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *_Service) Login(payload *entity.Login) (*entity.LoginResponse, error) {

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	respUser, err := s.repo.AuthRepo.Checklogin(payload)
	if err != nil {
		log.Println("check login is error : ", err)
		return nil, err
	}

	if ok := hashing.CheckPasswordHash(payload.Password, respUser.Password); !ok {
		log.Println("check password hashing is : ", ok)

		return nil, entity.CustomError("password is not match!")
	}

	expFromEnv := os.Getenv("EXPIRED")
	expire, _ := strconv.Atoi(expFromEnv)
	expirationTime := time.Now().Add(time.Duration(expire) * time.Hour)

	claims := &entity.Claims{
		ID:       respUser.ID,
		Username: respUser.FullName,
		Email:    respUser.Email,
		Role:     respUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(entity.JWTKEY)
	if err != nil {
		log.Println("Signed String error : ", err)
		return nil, err
	}

	return &entity.LoginResponse{FullName: respUser.FullName, Email: respUser.Email, Token: tokenString}, nil
}
