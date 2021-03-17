package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: no record was found")
	ErrDuplicateUsername  = errors.New("models: duplicate player username")
	ErrDuplicateCharacter = errors.New("models: duplicate character name")
)

// Player is the code representation of the "Player" relation in the
// database schema.
type Player struct {
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	Name     string `json:"name" db:"name"`
}

// Character is the code representation of the "Character" relation in
// the database schema.
type Character struct {
	Name           string `json:"name" db:"name"`
	Weight         int    `json:"weight" db:"weight"`
	Height         int    `json:"height" db:"height"`
	Alignment      string `json:"alignment" db:"alignment"`
	Sex            string `json:"sex" db:"sex"`
	Background     string `json:"background" db:"background"`
	Race           string `json:"race" db:"race"`
	Speed          int    `json:"speed" db:"speed"`
	Strength       int    `json:"strength" db:"strength"`
	Dexterity      int    `json:"dexterity" db:"dexterity"`
	Intelligence   int    `json:"intelligence" db:"intelligence"`
	Wisdom         int    `json:"wisdom" db:"wisdom"`
	Charisma       int    `json:"charisma" db:"charisma"`
	Constitution   int    `json:"constitution" db:"constitution"`
	HPMax          int    `json:"hp_max" db:"hp_max"`
	AbilityPoints  int    `json:"ability_points" db:"ability_points"`
	XPPoints       int    `json:"xp_points" db:"xp_points"`
	Class          string `json:"class" db:"class"`
	PlayerUsername string `json:"player_username" db:"player_username"`
}
