package repository

import (
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *CategoryRepositoryDB) UpdateCategoryByID(category model.Category) error {
	query := `
    UPDATE categories
    SET name = $1, description = $2
    WHERE id = $3`

	_, err := r.DB.Exec(query, category.Name, category.Description, category.ID)
	return err
}
