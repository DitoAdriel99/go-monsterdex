package auth

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

type Contract interface {
	Login(payload *entity.Login) (*entity.LoginResponse, error)
	Register(payload *entity.RegisterPayload) (*entity.Register, error)
}

type _Service struct {
	cfg   config.Cfg
	repo  *repository.Repo
	token tokenizer.JWT
}

func NewAuthService(cfg config.Cfg, repo *repository.Repo, token tokenizer.JWT) Contract {
	return &_Service{cfg, repo, token}
}
