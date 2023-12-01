package service

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/auth"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/monster"
)

type Service struct {
	MonsterService monster.Contract
	AuthService    auth.Contract
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		MonsterService: monster.NewService(repo),
		AuthService:    auth.NewAuthService(repo),
	}
}
