package monster

import (
	"fmt"
	"log"
	"time"
)

func (r *_Repo) SetStatus(monsterID int, status bool) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	queryUpdateMons := `
		UPDATE 
			monsters 
		SET 
			is_active = $2,
			updated_at = $3
		WHERE
			id = $1
	`

	if _, err := tx.Exec(queryUpdateMons, monsterID, status, time.Now().Local()); err != nil {
		log.Printf("error execute query set status: %v", err)
		return err
	}

	return nil
}
