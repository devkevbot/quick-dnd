package postgresql

import (
	"database/sql"
	"draco/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type CampaignModel struct {
	DB *sqlx.DB
}

func (m *CampaignModel) Insert(c models.Campaign) error {
	stmt := `INSERT INTO Campaign (name, current_location, state, dungeon_master)
		VALUES($1, $2, $3, $4)
		RETURNING id`

	var createdCampaignID int
	err := m.DB.QueryRowx(
		stmt, c.Name, c.CurrentLocation, c.State, c.DungeonMaster,
	).Scan(&createdCampaignID)

	return err
}

// Get attempts to retrieve a campaign entity from the database with an
// id equal to `id`.
func (m *CampaignModel) Get(id int) (*models.Campaign, error) {
	var storedCampaign models.Campaign

	stmt := `SELECT *
			FROM Campaign
			WHERE id = $1`
	row := m.DB.QueryRowx(stmt, id)

	if err := row.StructScan(&storedCampaign); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedCampaign, nil
}

// GetAllCharacterCampaigns retrieves all campaigns that `character` is a part of.
func (m *CampaignModel) GetAllCharacterCampaigns(characterID int) (*[]models.Campaign, error) {
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
