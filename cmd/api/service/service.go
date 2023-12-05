package service

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/auth"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service/monster"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	MonsterService monster.Contract
	AuthService    auth.Contract
}

func NewService(
	cfg config.Cfg,
	repo *repository.Repo,
	rdb *redis.Client,
	gcs storage.Storage,
	token tokenizer.JWT) *Service {
	return &Service{
		MonsterService: monster.NewService(cfg, repo, rdb, gcs, token),
		AuthService:    auth.NewAuthService(cfg, repo, token),
	}
}
