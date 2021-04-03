package postgresql

import (
	"database/sql"
	"draco/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type StatsModel struct {
	DB *sqlx.DB
}

// Retrieve all statistics
func (m *StatsModel) GetAll() (*models.Stats, error) {
	var storedStats models.Stats

	stmt := "SELECT * FROM Stats"
	row := m.DB.QueryRowx(stmt)

	if err := row.StructScan(&storedStats); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedStats, nil
}
