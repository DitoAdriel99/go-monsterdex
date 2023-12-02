package repository

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/auth"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/monster"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/monster_category"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/monster_type"
)

type Repo struct {
	MonsterRepo         monster.Contract
	MonsterTypeRepo     monster_type.Contract
	MonsterCategoryRepo monster_category.Contract
	AuthRepo            auth.Contract
}

func NewRepo() *Repo {
	return &Repo{
		MonsterRepo:         monster.NewRepositories(),
		MonsterTypeRepo:     monster_type.NewRepositories(),
		MonsterCategoryRepo: monster_category.NewRepositories(),
		AuthRepo:            auth.NewRepositories(),
	}
}
