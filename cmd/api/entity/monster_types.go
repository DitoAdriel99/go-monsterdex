package entity

type MonsterTypes []MonsterType

type MonsterType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
