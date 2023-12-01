package monster

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
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
	repo *repository.Repo
}

func NewService(repo *repository.Repo) Contract {
	return &_Service{repo}
}
