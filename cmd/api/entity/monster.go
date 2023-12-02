package entity

import (
	"fmt"
	"time"

	"github.com/DitoAdriel99/go-monsterdex/pkg/errbank"
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	MonsterRedisKey errbank.Error = "monster-"
)

var (
	ErrAlreadyActive   errbank.Error = "This monster is already Active!"
	ErrAlreadyDeactive errbank.Error = "This monster is already Deactive!"
)

type Monsters []Monster

type Monster struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	MonsterCategoryID int       `json:"monster_category_id"`
	Description       string    `json:"description,omitempty"`
	Image             string    `json:"image,omitempty"`
	IsCatched         bool      `json:"is_catched"`
	TypesID           []int     `json:"types_id,omitempty"`
	Height            float64   `json:"height,omitempty"`
	Weight            float64   `json:"weight,omitempty"`
	StatsHP           int       `json:"stats_hp,omitempty"`
	StatsAttack       int       `json:"stats_attack,omitempty"`
	StatsDefense      int       `json:"stats_defense,omitempty"`
	StatsSpeed        int       `json:"stats_speed,omitempty"`
	IsActive          bool      `json:"is_active,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type MonsterPayload struct {
	Name              string  `json:"name"`
	MonsterCategoryID int     `json:"monster_category_id"`
	Description       string  `json:"description"`
	Image             string  `json:"image"`
	TypesID           []int   `json:"types_id,omitempty"`
	Height            float64 `json:"height"`
	Weight            float64 `json:"weight"`
	StatsHP           int     `json:"stats_hp"`
	StatsAttack       int     `json:"stats_attack"`
	StatsDefense      int     `json:"stats_defense"`
	StatsSpeed        int     `json:"stats_speed"`
}

type StatusPayload struct {
	Status bool `json:"status"`
}

type Catch struct {
	UserID    int       `json:"user_id"`
	MonsterID int       `json:"monster_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type CatchPayload struct {
	Catch bool `json:"catch"`
}

func (m MonsterPayload) ParseIntArrayString() []int {
	var integers []int
	integers = append(integers, m.TypesID...)

	return integers
}

func (m StatusPayload) Validate() error {
	if m.Status != true && m.Status != false {
		return validation.Errors{"status": fmt.Errorf("Status must be a boolean value!")}
	}
	return nil
}

func (m MonsterPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&m.MonsterCategoryID, validation.Required),
		validation.Field(&m.Description, validation.Required, validation.Length(1, 255)),
		validation.Field(&m.Image, validation.Required),
		validation.Field(&m.TypesID, validation.Required, validation.Length(1, 0).Error("Type must contain at least one element")),
		validation.Field(&m.Height, validation.Required, validation.Min(0.0)),
		validation.Field(&m.Weight, validation.Required, validation.Min(0.0)),
		validation.Field(&m.StatsHP, validation.Required, validation.Min(0)),
		validation.Field(&m.StatsAttack, validation.Required, validation.Min(0)),
		validation.Field(&m.StatsDefense, validation.Required, validation.Min(0)),
		validation.Field(&m.StatsSpeed, validation.Required, validation.Min(0)),
	)
}
