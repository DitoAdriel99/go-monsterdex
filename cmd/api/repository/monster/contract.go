package monster

import (
	"database/sql"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
)

type _Repo struct {
	conn *sql.DB
}

type Contract interface {
	Create(userID int, req *entity.Monster) (int, error)
	Get(userID int, m *meta.Metadata) (*presentation.Monsters, error)
	GetID(userID int, monsterID int) (*presentation.Monster, error)
	GetIDAll(userID int, monsterID int) (*presentation.Monster, error)
	Update(monsterID int, req *entity.Monster) error
	SetStatus(monsterID int, status bool) error
	Catch(req *entity.Catch) (*bool, error)
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
