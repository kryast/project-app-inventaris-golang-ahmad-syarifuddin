package repository

import (
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *ItemRepositoryDB) GetAllItem() ([]model.Item, error) {
	query := `SELECT 
                  items.id,
                  items.name,
                  items.category_id,
                  categories.name AS category_name,
                  items.photo_url,
                  items.price,
                  items.purchase_date,
                  items.total_usage_days,
                  items.replacement_required
              FROM items
              INNER JOIN categories ON items.category_id = categories.id`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
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
			&item.ReplacementRequired,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
