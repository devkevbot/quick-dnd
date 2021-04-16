package postgresql

import "github.com/jmoiron/sqlx"

type MilestoneModel struct {
	DB *sqlx.DB
}

// Insert creates a new milestone corresponding with a campaign
// identified by `campaignID`.
func (m *MilestoneModel) Insert(campaignID int, milestone string) error {
	stmt := "INSERT INTO CampaignMilestones (campaign_id, milestone) VALUES ($1, $2)"

	_, err := m.DB.Exec(stmt, campaignID, milestone)
	if err != nil {
		return err
	}

	return nil
}

// GetAllForCampaign retrieves all milestone belonging to a campaign
// identified by `campaignID`.
func (m *MilestoneModel) GetAllForCampaign(campaignID int) (*[]string, error) {
	var storedMilestones []string

	stmt := "SELECT milestone FROM CampaignMilestones WHERE campaign_id = $1"

	rows, err := m.DB.Queryx(stmt, campaignID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var milestone string
		err = rows.Scan(&milestone)
		if err != nil {
			return nil, err
		}
		storedMilestones = append(storedMilestones, milestone)
	}

	return &storedMilestones, nil
}
