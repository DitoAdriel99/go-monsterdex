package monster

import (
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/lib/pq"
)

func (r *_Repo) Update(monsterID int, req *entity.Monster) error {
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
			name = $1, 
			monster_category_id = $2, 
			description= $3, 
			image = $4,
			types_id = $5, 
			height = $6, 
			weight = $7, 
			stats_hp = $8, 
			stats_attack = $9, 
			stats_defense = $10, 
			stats_speed = $11,
			updated_at = $12
		WHERE
			id = $13
	`

	typeArray := pq.Array(req.TypesID)

	if _, err := tx.Exec(queryUpdateMons, req.Name, req.MonsterCategoryID, req.Description, req.Image, typeArray, req.Height, req.Weight, req.StatsHP, req.StatsAttack, req.StatsDefense, req.StatsSpeed, req.UpdatedAt, monsterID); err != nil {
		log.Printf("error execute query update: %v", err)
		return err
	}

	return nil
}
