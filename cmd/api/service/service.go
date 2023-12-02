package service

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/auth"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/monster"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	MonsterService monster.Contract
	AuthService    auth.Contract
}

func NewService(repo *repository.Repo, rdb *redis.Client) *Service {
	return &Service{
		MonsterService: monster.NewService(repo, rdb),
		AuthService:    auth.NewAuthService(repo),
	}
}
