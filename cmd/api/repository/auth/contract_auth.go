package auth

import (
	"database/sql"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"

	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/google/uuid"
)

type _Repo struct {
	conn *sql.DB
}

type Contract interface {
	Checklogin(auth *entity.Login) (*entity.User, error)
	ValidateUser(email string) (*entity.User, error)
	CheckEmail(email string) error
	Register(rq *entity.Register) (int, error)
	ValidateUserId(id uuid.UUID) (*entity.User, error)
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
