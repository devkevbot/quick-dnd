package postgresql

import (
	"draco/models"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type BelongsToModel struct {
	DB *sqlx.DB
}

func (m *BelongsToModel) Insert(c models.BelongsTo) error {
	stmt := `INSERT INTO BelongsTo (character_id, campaign_id)
		VALUES($1, $2)`

	_, err := m.DB.Exec(stmt, c.CharacterID, c.CampaignID)
	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if postgresError.Code.Name() == "unique_violation" {
				return models.ErrDuplicateBelongsTo
			}
		}
		return err
	}
	return nil
}

// Retrieves all campaigns that `character` is a part of.
func (m *BelongsToModel) GetAllCharacterCampaigns(characterID int) (*[]models.Campaign, error) {
	var storedCampaigns []models.Campaign

	stmt := `SELECT *
			FROM Campaign
			WHERE id IN (
				SELECT campaign_id
				FROM BelongsTo
				WHERE character_id = $1
		)`
	rows, err := m.DB.Queryx(stmt, characterID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Campaign
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		storedCampaigns = append(storedCampaigns, c)
	}

	return &storedCampaigns, nil
}

// Retrieves all characters that are part of `campaign`
func (m *BelongsToModel) GetAllCampaignCharacters(campaignID int) (*[]models.Character, error) {
	var storedCharacters []models.Character

	stmt := `SELECT *
			FROM Character
			WHERE id IN (
				SELECT character_id
				FROM BelongsTo
				WHERE campaign_id = $1
		)`
	rows, err := m.DB.Queryx(stmt, campaignID)
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
