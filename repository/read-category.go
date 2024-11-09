package repository

import "project-app-inventaris-golang-ahmad-syarifuddin/model"

func (r *CategoryRepositoryDB) GetAllCategory() ([]model.Category, error) {
	categories := []model.Category{}
	query := `SELECT id, name, description FROM categories`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category model.Category
		rows.Scan(&category.ID, &category.Name, &category.Description)

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
