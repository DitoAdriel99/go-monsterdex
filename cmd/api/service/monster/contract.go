package monster

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
	"github.com/DitoAdriel99/go-monsterdex/pkg/storage"
	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"
	"github.com/go-redis/redis/v8"
)

type Contract interface {
	Create(req *entity.MonsterPayload) (*entity.Monster, error)
	Get(bearer string, m *meta.Metadata) (*presentation.Monsters, error)
	GetID(bearer string, monsterID int) (*presentation.Monster, error)
	Update(monsterID int, req *entity.MonsterPayload) (*entity.Monster, error)
	SetStatus(monsterID int, req *entity.StatusPayload) error
	Catch(bearer string, monsterID int) (*bool, error)
}

type _Service struct {
	cfg   config.Cfg
	repo  *repository.Repo
	rdb   *redis.Client
	gcs   storage.Storage
	token tokenizer.JWT
}

func NewService(cfg config.Cfg, repo *repository.Repo, rdb *redis.Client, gcs storage.Storage, token tokenizer.JWT) Contract {
	return &_Service{cfg, repo, rdb, gcs, token}
}
