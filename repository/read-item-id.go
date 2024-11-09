package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *ItemRepositoryDB) GetItemById(id int) (*model.Item, error) {

	query := `SELECT 
                  items.id,
                  items.name,
                  items.category_id,
                  categories.name AS category_name,  -- Menambahkan category_name
                  items.photo_url,
                  items.price,
                  items.purchase_date,
                  items.total_usage_days
              FROM items
              INNER JOIN categories ON items.category_id = categories.id
              WHERE items.id = $1`

	rows := r.DB.QueryRow(query, id)

	var item model.Item

	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.CategoryID,
		&item.CategoryName,
		&item.PhotoURL,
		&item.Price,
		&item.PurchaseDate,
		&item.TotalUsageDay,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item with id %d not found", id)
		}
		return nil, err
	}

	return &item, nil
}
