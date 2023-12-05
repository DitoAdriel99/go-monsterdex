package auth

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/pkg/hashing"

	"log"
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

	tokenString, err := s.token.GenerateToken(*respUser)
	if err != nil {
		log.Println("Signed String error : ", err)
		return nil, err
	}

	return &entity.LoginResponse{FullName: respUser.FullName, Email: respUser.Email, Token: tokenString}, nil
}
