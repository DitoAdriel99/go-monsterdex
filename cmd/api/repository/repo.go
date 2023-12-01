package repository

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/auth"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository/monster"
)

type Repo struct {
	MonsterRepo monster.Contract
	AuthRepo    auth.Contract
}

func NewRepo() *Repo {
	return &Repo{
		MonsterRepo: monster.NewRepositories(),
		AuthRepo:    auth.NewRepositories(),
	}
}
