package postgresql

import (
	"database/sql"
	"draco/models"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CharacterModel struct {
	DB *sqlx.DB
}

func (m *CharacterModel) Insert(
	name string, weight, height int,
	alignment, sex, background, race string,
	speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
	hpMax, abilityPoints, xpPoints int,
	class, playerUsername string,
) error {
	var err error

	stmt := `INSERT INTO Character (name, weight, height,
		alignment, sex, background, race,
		speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
		hp_max, ability_points, xp_points,
		class, player_username)
	    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`

	_, err = m.DB.Exec(
		stmt, name, weight, height,
		alignment, sex, background, race,
		speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
		hpMax, abilityPoints, xpPoints,
		class, playerUsername,
	)

	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if postgresError.Code.Name() == "unique_violation" {
				if strings.Contains(postgresError.Message, "character_pkey") {
					return models.ErrDuplicateCharacter
				}
			}
		}
		return err
	}

	return nil
}

// Get attempts to retrieve a character entity from the database with a
// name equal to `name`.
func (m *CharacterModel) Get(name string) (*models.Character, error) {
	var storedCharacter models.Character

	stmt := "SELECT * FROM Character WHERE name = $1"
	row := m.DB.QueryRowx(stmt, name)

	if err := row.StructScan(&storedCharacter); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedCharacter, nil
}
