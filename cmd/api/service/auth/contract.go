package auth

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

type Contract interface {
	Login(payload *entity.Login) (*entity.LoginResponse, error)
	Register(payload *entity.RegisterPayload) (*entity.Register, error)
}

type _Service struct {
	repo *repository.Repo
}

func NewAuthService(repo *repository.Repo) Contract {
	return &_Service{repo}
}
