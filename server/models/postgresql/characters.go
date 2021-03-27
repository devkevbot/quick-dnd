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
	alignment models.AlignmentType, sex models.SexType, background string, race models.RaceType,
	speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
	hpMax, abilityPoints, xpPoints int,
	class models.ClassType, playerUsername string,
) (int, error) {
	stmt := `INSERT INTO Character (name, weight, height,
		alignment, sex, background, race,
		speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
		hp_max, ability_points, xp_points,
		class, player_username)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id`

	var createdCharacterID int
	err := m.DB.QueryRowx(
		stmt, name, weight, height,
		alignment, sex, background, race,
		speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
		hpMax, abilityPoints, xpPoints,
		class, playerUsername,
	).Scan(&createdCharacterID)

	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if postgresError.Code.Name() == "unique_violation" {
				if strings.Contains(postgresError.Message, "character_name_player_username_key") {
					return -1, models.ErrDuplicateCharacter
				}
			}
		}
		return -1, err
	}

	return createdCharacterID, nil
}

// Get attempts to retrieve a character entity from the database with a
// name equal to `name`.
func (m *CharacterModel) Get(id int) (*models.Character, error) {
	var storedCharacter models.Character

	stmt := "SELECT * FROM Character WHERE id = $1"
	row := m.DB.QueryRowx(stmt, id)

	if err := row.StructScan(&storedCharacter); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedCharacter, nil
}
