package monster_category

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

func (r *_Repo) GetId(monsterID int) (*entity.MonsterCategory, error) {
	query := `
		SELECT 
			id, name
		FROM 
			monster_types
		WHERE
			id = $1;
	`

	var monsters entity.MonsterCategory
	err := r.conn.QueryRow(query, monsterID).Scan(&monsters.ID, &monsters.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.CustomError("Monster Category Not Found!")
		}
		log.Println("error executing query select by id to monster", err)
		return nil, fmt.Errorf("query execution error: %w", err)
	}

	return &monsters, nil
}

// Helper function to repeat a string
func repeatStr(str string, times int) []string {
	var result []string
	count := 0

	for i := 0; i < times; i++ {
		count++
		newStr := fmt.Sprintf("%s%d", str, count)
		result = append(result, newStr)
	}
	return result
}
