package repository

func (cr *CategoryRepositoryDB) DeleteCategoryByID(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := cr.DB.Exec(query, id)
	return err
}
