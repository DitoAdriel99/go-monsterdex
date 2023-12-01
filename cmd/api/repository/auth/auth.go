package auth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/google/uuid"
)

func (c *_Repo) Checklogin(auth *entity.Login) (*entity.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var object entity.User

	err := c.conn.QueryRow(query, auth.Email).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.CustomError("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil

}

func (c *_Repo) ValidateUser(email string) (*entity.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var object entity.User

	err := c.conn.QueryRow(query, email).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.CustomError("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	return &object, nil
}

func (c *_Repo) CheckEmail(email string) error {
	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	var count int
	err := c.conn.QueryRow(query, email).Scan(&count)
	if err != nil {
		log.Println("Error querying select users", err)
		err = fmt.Errorf("scanning activity objects: %w", err)
		return err
	}

	if count == 1 {
		err = entity.CustomError("Email Already Used!")
		return err
	}

	return nil
}

func (c *_Repo) Register(rq *entity.Register) (int, error) {
	queryInsert := `
		INSERT INTO users (fullname, email, password, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`

	var id int
	if err := c.conn.QueryRow(queryInsert, rq.FullName, rq.Email, rq.Password, rq.Role, rq.CreatedAt, rq.UpdatedAt).Scan(&id); err != nil {
		log.Println("error executing query insert to users", err)

		err = fmt.Errorf("executing query: %w", err)
		return 0, err
	}

	return id, nil
}

func (c *_Repo) ValidateUserId(id uuid.UUID) (*entity.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var object entity.User

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.FullName,
		&object.Email,
		&object.Password,
		&object.Role,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.CustomError("user not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	return &object, nil
}
