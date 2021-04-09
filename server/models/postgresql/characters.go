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

func (m *CharacterModel) Insert(c models.Character) (int, error) {
	stmt := `INSERT INTO Character (name, weight, height,
		alignment, sex, background, race,
		speed, strength, dexterity, intelligence, wisdom, charisma, constitution,
		hp_max, ability_points, xp_points,
		class, player_username)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id`

	var createdCharacterID int
	err := m.DB.QueryRowx(
		stmt, c.Name, c.Weight, c.Height,
		c.Alignment, c.Sex, c.Background, c.Race,
		c.Speed, c.Strength, c.Dexterity, c.Intelligence, c.Wisdom, c.Charisma, c.Constitution,
		c.HPMax, c.AbilityPoints, c.XPPoints,
		c.Class, c.PlayerUsername,
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

// GetAllUserCharacters retrieves all characters that belong to `username`.
func (m *CharacterModel) GetAllUserCharacters(username string) (*[]models.Character, error) {
	var storedCharacters []models.Character

	stmt := "SELECT * FROM Character WHERE player_username = $1"
	rows, err := m.DB.Queryx(stmt, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Character
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		storedCharacters = append(storedCharacters, c)
	}

	return &storedCharacters, nil
}

// Delete attempts to delete a character identified by `id`.
func (m *CharacterModel) Delete(id int) error {
	stmt := "DELETE FROM Character WHERE id = $1"

	res, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return models.ErrDeleteSingleRecord
	}

	return nil
}
