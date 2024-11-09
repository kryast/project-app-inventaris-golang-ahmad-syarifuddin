package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

type ItemRepositoryDB struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepositoryDB {
	return ItemRepositoryDB{DB: db}
}

func (r *ItemRepositoryDB) CreateDataItem(item model.Item) error {
	query := "INSERT INTO items (name , category_id, photo_url, price, purchase_date, total_usage_days, replacement_required) VALUES ($1 , $2 , $3, $4, $5, $6, $7) RETURNING id"

	err := r.DB.QueryRow(query, item.Name, item.CategoryID, item.PhotoURL, item.Price, item.PurchaseDate, item.TotalUsageDay, item.ReplacementRequired).Scan(&item.ID)
	if err != nil {
		return fmt.Errorf("error at repository")
	}

	return nil
}
