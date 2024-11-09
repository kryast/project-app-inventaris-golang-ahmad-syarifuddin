package repository

import (
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *ItemRepositoryDB) UpdateItemByID(item model.Item) error {
	query := `
    UPDATE items
    SET name = $1, category_id = $2, photo_url = $3, price = $4, purchase_date = $5
    WHERE id = $6`

	_, err := r.DB.Exec(query, item.Name, item.CategoryID, item.PhotoURL, item.Price, item.PurchaseDate, item.ID)
	return err
}
