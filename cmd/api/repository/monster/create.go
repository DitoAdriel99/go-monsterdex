package monster

import (
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/lib/pq"
)

func (r *_Repo) Create(userID int, req *entity.Monster) (int, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	queryInsertMons := `
		INSERT INTO monsters
			(name, monster_category_id, description, image, types_id, height, weight,  stats_hp, stats_attack, stats_defense, stats_speed, is_active, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id
	`

	var id int
	typeArray := pq.Array(req.TypesID)
	if err := tx.QueryRow(queryInsertMons, req.Name, req.MonsterCategoryID, req.Description, req.Image, typeArray, req.Height, req.Weight, req.StatsHP, req.StatsAttack, req.StatsDefense, req.StatsSpeed, req.IsActive, req.CreatedAt, req.UpdatedAt).Scan(&id); err != nil {
		log.Println("error executing query insert monster err:", err)
		return 0, err
	}

	return id, nil
}
