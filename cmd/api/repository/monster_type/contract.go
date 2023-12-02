package monster_type

import (
	"database/sql"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/config"
)

type _Repo struct {
	conn *sql.DB
}

type Contract interface {
	GetIds(monsterIDs []int) ([]entity.MonsterType, error)
}

func NewRepositories() Contract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_Repo{
		conn: conn,
	}
}

func (r *_Repo) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name", "id":
		return true
	default:
		return false
	}
}
