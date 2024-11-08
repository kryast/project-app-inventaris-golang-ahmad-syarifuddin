package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *CategoryRepositoryDB) GetCategoryById(id int) (*model.Category, error) {
	query := `SELECT id, name, description FROM categories where id = $1`
	rows := r.DB.QueryRow(query, id)

	var category model.Category
	err := rows.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	// Mengembalikan data pengguna
	return &category, nil
}
