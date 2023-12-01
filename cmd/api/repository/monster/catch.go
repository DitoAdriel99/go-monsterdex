package monster

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

func (r *_Repo) Catch(req *entity.Catch) (*bool, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	checkCatchQuery := `
		SELECT 
			is_catched 
		FROM 
			monster_owner 
		WHERE 
			user_id = $1 
		AND 
			monster_id = $2
	`

	var isCatched bool
	if err := tx.QueryRow(checkCatchQuery, req.UserID, req.MonsterID).Scan(&isCatched); err != nil {
		if err == sql.ErrNoRows {
			newStat := false
			queryInsertMons := `
				INSERT INTO monster_owner
					(monster_id, user_id, is_catched, created_at, updated_at) 
				VALUES 
					($1, $2, $3, $4, $5)
			`

			if _, err := tx.Exec(queryInsertMons, req.MonsterID, req.UserID, true, req.CreatedAt, req.UpdatedAt); err != nil {
				log.Println("error executing query insert to monster owner:", err)
				return nil, err
			}
			return &newStat, nil
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	var queryUpdate string

	if isCatched {
		queryUpdate = `
			UPDATE 
				monster_owner
			SET
				is_catched = false
			WHERE
				user_id = $1
			AND
				monster_id = $2
		`
	} else {
		queryUpdate = `
		UPDATE 
			monster_owner
		SET
			is_catched = true
		WHERE
			user_id = $1
		AND
			monster_id = $2
	`
	}
	if _, err := tx.Exec(queryUpdate, req.UserID, req.MonsterID); err != nil {
		log.Println("error executing query catch:", err)
		return nil, err
	}

	return &isCatched, nil
}
