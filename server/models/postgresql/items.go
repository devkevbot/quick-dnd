package postgresql

import (
	"database/sql"
	"draco/models"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ItemModel struct {
	DB *sqlx.DB
}

// Insert attempts to insert `i` into the Items table.
func (m *ItemModel) Insert(i models.Item) error {
	stmt := `INSERT INTO Items
	(character_id, item_name, type,
	rarity, weight, gold_value,
	quantity, description)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := m.DB.Exec(stmt,
		i.CharacterID, i.ItemName, i.Type,
		i.Rarity, i.Weight, i.GoldValue,
		i.Quantity, i.Description)
	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if postgresError.Code.Name() == "unique_violation" {
				return models.ErrDuplicateItem
			}
		}
		return err
	}

	return nil
}

// Get attempts to retrieve an item identified by `characterID` and `itemName`.
func (m *ItemModel) Get(characterID int, itemName string) (*models.Item, error) {
	var storedItem models.Item

	stmt := "SELECT * FROM Items WHERE character_id = $1 AND item_name = $2"
	row := m.DB.QueryRowx(stmt, characterID, itemName)

	if err := row.StructScan(&storedItem); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &storedItem, nil
}

// GetAllCharacterItems retrieves all items belonging to a character
// with `characterID`.
func (m *ItemModel) GetAllCharacterItems(characterID int) (*[]models.Item, error) {
	var storedItems []models.Item

	stmt := "SELECT * FROM Items WHERE character_id = $1"
	rows, err := m.DB.Queryx(stmt, characterID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i models.Item
		err = rows.StructScan(&i)
		if err != nil {
			return nil, err
		}
		storedItems = append(storedItems, i)
	}

	return &storedItems, nil
}

func (m *ItemModel) Delete(characterID int, itemName string) error {
	stmt := "DELETE FROM Items WHERE character_id = $1 AND item_name = $2"

	res, err := m.DB.Exec(stmt, characterID, itemName)
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

func (m *ItemModel) GetTotalWeightCharacterItems(characterID int) (*int, error) {
	var weight int

	stmt := "SELECT SUM(weight*quantity) FROM Items WHERE character_id = $1"
	row := m.DB.QueryRowx(stmt, characterID)

	err := row.Scan(&weight)
	if err != nil {
		return nil, err
	}
	return &weight, nil
}
