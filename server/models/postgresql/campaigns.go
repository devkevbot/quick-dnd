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

// Update attempts to update the state and location of the campaign
// identified by `id`.
func (m *CampaignModel) Update(id int, state string, location string) error {
	stmt := "UPDATE Campaign SET state = $2, current_location = $3 WHERE id = $1"

	res, err := m.DB.Exec(stmt, id, state, location)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return models.ErrUpdateSingleRecord
	}

	return nil

}

// Delete attempts to delete a campaign identified by `id`.
func (m *CampaignModel) Delete(id int) error {
	stmt := "DELETE FROM Campaign WHERE id = $1"

	res, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return models.ErrNoRecord

	}
	if count > 1 {
		return models.ErrDeleteSingleRecord
	}

	return nil
}

// Get all campaigns started by `dungeonMaster`.
func (m *CampaignModel) GetPlayersCreatedCampaigns(dungeonMaster string) (*[]models.Campaign, error) {
	var storedCampaigns []models.Campaign

	stmt := `SELECT *
			FROM Campaign
			WHERE dungeon_master = $1`

	rows, err := m.DB.Queryx(stmt, dungeonMaster)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var campaign models.Campaign
		err = rows.StructScan(&campaign)
		if err != nil {
			return nil, err
		}
		storedCampaigns = append(storedCampaigns, campaign)
	}

	return &storedCampaigns, nil
}

func (m *CampaignModel) GetAllCharacterCampaigns(characterID int) (*[]models.Campaign, error) {
	var storedCampaigns []models.Campaign

	stmt := `SELECT *
			FROM Campaign
			WHERE id IN (
				SELECT campaign_id
				FROM BelongsTO
				WHERE character_id = $1
			)`

	rows, err := m.DB.Queryx(stmt, characterID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var campaign models.Campaign
		err = rows.StructScan(&campaign)
		if err != nil {
			return nil, err
		}
		storedCampaigns = append(storedCampaigns, campaign)
	}

	return &storedCampaigns, nil
}

// GetCampaignParticpants fetches some data about the players and
// characters that belong to a campaign identified by `id`.
func (m *CampaignModel) GetCampaignParticpants(id int) (*[]models.CampaignParticipants, error) {
	var participants []models.CampaignParticipants

	stmt := `SELECT ch.player_username, ch.name, ch.race, ch.class
			FROM Character AS ch
			INNER JOIN BelongsTo
			ON BelongsTo.character_id = ch.id
			INNER JOIN Campaign AS ca
			ON BelongsTo.campaign_id = ca.id
			WHERE ca.id = $1`

	rows, err := m.DB.Queryx(stmt, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var participant models.CampaignParticipants
		err = rows.Scan(
			&participant.PlayerUsername,
			&participant.CharacterName,
			&participant.CharacterRace,
			&participant.CharacterClass,
		)
		if err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return &participants, nil
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
