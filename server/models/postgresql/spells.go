package postgresql

import (
	"database/sql"
	"draco/models"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SpellModel struct {
	DB *sqlx.DB
}

// Insert attempts to insert `s` into the Spells table.
func (m *SpellModel) Insert(s models.Spell) error {
	stmt := `INSERT INTO SPELLS
	(character_id, spell_name, level, school, concentration,
	description, casting_time, range, duration)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.Exec(stmt,
		s.CharacterID, s.SpellName, s.Level, s.School, s.Concentration,
		s.Description, s.CastingTime, s.Range, s.Duration)
	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if postgresError.Code.Name() == "unique_violation" {
				return models.ErrDuplicateSpell
			}
		}
		return err
	}

	return nil
}

// Get attempts to retrieve a spell identified by `characterID` and `spellName`.
func (m *SpellModel) Get(characterID int, spellName string) (*models.Spell, error) {
	var storedSpell models.Spell

	stmt := "SELECT * FROM Spells WHERE character_id = $1 AND spell_name = $2"
	row := m.DB.QueryRowx(stmt, characterID, spellName)

	if err := row.StructScan(&storedSpell); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedSpell, nil
}

// GetAllCharacterSpells retrieves all spells belonging to a character
// with `characterID`.
func (m *SpellModel) GetAllCharacterSpells(characterID int) (*[]models.Spell, error) {
	var storedSpells []models.Spell

	stmt := "SELECT * FROM Spells WHERE character_id = $1"
	rows, err := m.DB.Queryx(stmt, characterID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s models.Spell
		err = rows.StructScan(&s)
		if err != nil {
			return nil, err
		}
		storedSpells = append(storedSpells, s)
	}

	return &storedSpells, nil
}

// Delete a spell belonging to a character.
func (m *SpellModel) Delete(characterID int, spellName string) error {
	stmt := "DELETE FROM Spells WHERE character_id = $1 AND spell_name = $2"

	res, err := m.DB.Exec(stmt, characterID, spellName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoRecord
		} else {
			return err
		}
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
