package monster

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
)

func (r *_Repo) GetIDAll(userID int, monsterID int) (*presentation.Monster, error) {
	stmt := `SELECT 
                mn.id,
                mn.name,
                mc.id,
                mc.name,
                mn.description,
                mn.image,
                COALESCE(mo.is_catched, false) AS is_catched,
				array_agg(mt.id) AS type_id,
				array_agg(mt.name) AS type_names,
                mn.height,
                mn.weight,
                mn.stats_hp,
                mn.stats_attack,
                mn.stats_defense,
                mn.stats_speed,
                mn.is_active,
                mn.created_at,
                mn.updated_at
				FROM 
                monsters mn
            LEFT JOIN
                monster_owner mo
            ON
                mn.id = mo.monster_id 
					AND 
						mo.user_id = $1
			LEFT JOIN
				monster_types mt 
			ON 
				CAST(mt.id AS TEXT) = ANY(mn.types_id)
			LEFT JOIN
				monster_category mc
			ON
				mc.id = mn.monster_category_id
            WHERE 
                mn.id = $2
			GROUP BY
    			mn.id,
    			mn.name,
                mc.id,
        		mc.name,
				mn.types_id,
				mn.image,
				mo.is_catched,
				mn.is_active,
				mn.created_at,
				mn.updated_at
            `
	args := []interface{}{userID, monsterID}

	stmt += ";"

	var isCaught bool
	var typeStr string
	var typeInt string
	var collection presentation.Monster
	if err := r.conn.QueryRow(stmt, args...).Scan(
		&collection.ID,
		&collection.Name,
		&collection.MonsterCategory.ID,
		&collection.MonsterCategory.Name,
		&collection.Description,
		&collection.Image,
		&isCaught,
		&typeInt,
		&typeStr,
		&collection.Height,
		&collection.Weight,
		&collection.StatsHP,
		&collection.StatsAttack,
		&collection.StatsDefense,
		&collection.StatsSpeed,
		&collection.IsActive,
		&collection.CreatedAt,
		&collection.UpdatedAt,
	); err != nil {
		log.Println("error executing query select by id to monster", err)
		if err == sql.ErrNoRows {
			return nil, entity.CustomError("Monster Not Found!")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}
	if isCaught {
		collection.IsCatched = true
	} else {
		collection.IsCatched = false
	}

	// Parse string representations into slices of integers and strings
	parsedTypeIDs, err := parseIntArrayString(typeInt)
	if err != nil {
		return nil, err
	}

	parsedTypeNames := parseTypeString(typeStr)

	collection.Types = make([]presentation.MonsterType, len(parsedTypeIDs))

	for i, v := range parsedTypeIDs {
		collection.Types[i].ID = v
		collection.Types[i].Name = parsedTypeNames[i]
	}

	return &collection, nil
}
