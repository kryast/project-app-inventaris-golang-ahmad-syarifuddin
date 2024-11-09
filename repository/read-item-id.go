package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *ItemRepositoryDB) GetItemById(id int) (*model.Item, error) {
	query := `SELECT id, name, category_id, photo_url, price, purchase_date, total_usage_days FROM items where id = $1`
	rows := r.DB.QueryRow(query, id)

	var item model.Item
	err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.PhotoURL, &item.Price, &item.PurchaseDate, &item.TotalUsageDay)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &item, nil
}
