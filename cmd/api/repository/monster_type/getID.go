package monster_type

import (
	"fmt"
	"log"
	"strings"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
)

func (r *_Repo) GetIds(monsterIDs []int) ([]entity.MonsterType, error) {
	query := `
		SELECT 
			id, name
		FROM 
			monster_types
		WHERE
			id IN (` + strings.Join(repeatStr("$", len(monsterIDs)), ",") + `);
	`

	args := make([]interface{}, len(monsterIDs))
	for i, id := range monsterIDs {
		args[i] = id
	}

	rows, err := r.conn.Query(query, args...)
	if err != nil {
		log.Println("error executing query select by id to monster", err)
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	var monsters []entity.MonsterType
	for rows.Next() {
		var monster entity.MonsterType
		if err := rows.Scan(&monster.ID, &monster.Name); err != nil {
			log.Println("error scanning row", err)
			return nil, fmt.Errorf("scanning row error: %w", err)
		}
		monsters = append(monsters, monster)
	}

	if err := rows.Err(); err != nil {
		log.Println("error iterating rows", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(monsters) != len(monsterIDs) {
		return nil, entity.CustomError("Monster Types Not Found")
	}

	return monsters, nil
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
