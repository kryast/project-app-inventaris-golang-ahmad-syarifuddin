package repository

func (r *ItemRepositoryDB) DeleteItemByID(id int) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
