package monster

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/presentation"
	"github.com/DitoAdriel99/go-monsterdex/pkg/meta"
)

func (r *_Repo) Get(userID int, m *meta.Metadata) (*presentation.Monsters, error) {
	q, err := meta.ParseMetaData(m, r)
	if err != nil {
		return nil, err
	}
	stmt := `SELECT 
                mn.id,
                mn.name,
                mc.id,
                mc.name,
				mn.image,
                COALESCE(mo.is_catched, false) AS is_catched,
                array_agg(mt.id) AS type_id,
                array_agg(mt.name) AS type_names,
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
                mn.is_active = $2 
            `
	args := []interface{}{userID, true}
	// If q.Type has at least one value
	if len(q.Type) > 0 {
		stmt += " AND mt.name ILIKE ANY(ARRAY["
		for i, t := range q.Type {
			if i > 0 {
				stmt += ","
			}
			stmt += fmt.Sprintf("$%d", len(args)+1)
			args = append(args, t)
		}
		stmt += "])"
	}

	if q.Name != "" {
		stmt += " AND m.name ILIKE '%' || $" + strconv.Itoa(len(args)+1) + " || '%'"
		args = append(args, q.Name)
	}

	if len(q.IsCatched) > 0 {
		stmt += " AND mo.is_catched = $" + strconv.Itoa(len(args)+1)
		args = append(args, q.IsCatched)
	}

	// Check for sort order by ID

	stmt += ` GROUP BY
				mn.id,
				mn.name,
				mc.id,
				mc.name,
				mn.types_id,
				mn.image,
				mo.is_catched,
				mn.is_active,
				mn.created_at,
				mn.updated_at`

	switch q.OrderDirection {
	case "asc", "ascending":
		stmt += fmt.Sprintf(" ORDER BY mn.%s ASC", q.OrderBy)
	case "desc", "descending":
		stmt += fmt.Sprintf(" ORDER BY mn.%s DESC", q.OrderBy)
	default:
		stmt += " ORDER BY mn.id DESC"
	}

	stmt += ";"

	rows, err := r.conn.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make(presentation.Monsters, 0)
	for rows.Next() {
		var typeStr string
		var typeInt string
		var isCaught bool
		var collection presentation.Monster
		if err := rows.Scan(
			&collection.ID,
			&collection.Name,
			&collection.MonsterCategory.ID,
			&collection.MonsterCategory.Name,
			&collection.Image,
			&isCaught,
			&typeInt,
			&typeStr,
			&collection.IsActive,
			&collection.CreatedAt,
			&collection.UpdatedAt,
		); err != nil {
			log.Println("error executing query select to monster", err)
			return nil, err
		}

		if isCaught {
			collection.IsCatched = true
		} else {
			collection.IsCatched = false
		}

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
		collections = append(collections, collection)
	}
	return &collections, nil
}

func parseTypeString(arrayStr string) []string {
	arrayStr = strings.Trim(arrayStr, "{}")
	types := strings.Split(arrayStr, ",")

	for i := range types {
		types[i] = strings.TrimSpace(types[i])
	}

	return types
}

func parseIntArrayString(arrayStr string) ([]int, error) {
	arrayStr = strings.Trim(arrayStr, "{}")
	items := strings.Split(arrayStr, ",")

	var integers []int
	for _, item := range items {
		trimmedItem := strings.TrimSpace(item)
		intVal, err := strconv.Atoi(trimmedItem)
		if err != nil {
			return nil, err
		}
		integers = append(integers, intVal)
	}

	return integers, nil
}
