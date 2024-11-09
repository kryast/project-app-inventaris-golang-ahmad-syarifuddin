package repository

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

// Implementasi fungsi untuk mengambil item yang memerlukan penggantian
func (r *ItemRepositoryDB) GetItemsForReplacement() ([]model.Item, error) {
	query := `
        SELECT id, name, category_id, photo_url, price, purchase_date, total_usage_days, replacement_required
        FROM items
        WHERE replacement_required = true
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.PhotoURL, &item.Price, &item.PurchaseDate, &item.TotalUsageDay, &item.ReplacementRequired); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
