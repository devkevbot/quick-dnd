package models

import "errors"

var (
	ErrNoRecord          = errors.New("models: no record was found")
	ErrDuplicateUsername = errors.New("models: duplicate player username")
)

// Player is the code representation of the "Player" relation in the
// database schema.
type Player struct {
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	Name     string `json:"name" db:"name"`
}
