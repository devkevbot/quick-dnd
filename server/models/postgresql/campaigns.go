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

// Inserts a new campaign into the DB and any accompanying characters into the
// BelongsTo relation
func (m *CampaignModel) Insert(c models.Campaign, characterIDs []int) (int, error) {
	stmtCampaign := `INSERT INTO Campaign (name, current_location, state, dungeon_master)
		VALUES($1, $2, $3, $4)
		RETURNING id`
	stmtBelongsTo := `INSERT INTO BelongsTo (character_id, campaign_id)
	VALUES($1, $2)`

	var createdCampaignID int

	tx, err := m.DB.Beginx()
	if err != nil {
		return -1, err
	}

	err = tx.QueryRowx(
		stmtCampaign, c.Name, c.CurrentLocation, c.State, c.DungeonMaster,
	).Scan(&createdCampaignID)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	for _, id := range characterIDs {
		_, err := tx.Exec(stmtBelongsTo, id, createdCampaignID)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	err = tx.Commit()

	return createdCampaignID, err
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

func (m *CampaignModel) GetAllCharacterCampaigns(characterID int) {
	/* TODO: Implement this function */
}

// GetPlayersAttendedAll fetches the usernames of players which have
// participated in all campaigns which were created by `dungeonMaster`.
func (m *CampaignModel) GetPlayersAttendedAll(dungeonMaster string) (*[]string, error) {
	var storedUsernames []string

	stmt := `SELECT DISTINCT ch.player_username
			FROM Character ch
			WHERE NOT EXISTS
				(SELECT ca.id
				FROM Campaign ca
				WHERE ca.dungeon_master = $1 AND ca.id NOT IN
					(SELECT campaign_id
					FROM BelongsTo bt
					WHERE bt.character_id = ch.id))`

	rows, err := m.DB.Queryx(stmt, dungeonMaster)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		storedUsernames = append(storedUsernames, username)
	}

	return &storedUsernames, nil
}
