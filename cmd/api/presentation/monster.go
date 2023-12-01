package presentation

import "time"

type Monsters []Monster

type Monster struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	MonsterCategory MonsterCategory `json:"monster_category"`
	Description     string          `json:"description,omitempty"`
	Image           string          `json:"image,omitempty"`
	IsCatched       bool            `json:"is_catched"`
	Types           []MonsterType   `json:"monster_types,omitempty"`
	Height          float64         `json:"height,omitempty"`
	Weight          float64         `json:"weight,omitempty"`
	StatsHP         int             `json:"stats_hp,omitempty"`
	StatsAttack     int             `json:"stats_attack,omitempty"`
	StatsDefense    int             `json:"stats_defense,omitempty"`
	StatsSpeed      int             `json:"stats_speed,omitempty"`
	IsActive        bool            `json:"is_active,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type MonsterCategory struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MonsterType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
