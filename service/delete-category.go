package service

func (cs *CategoryService) DeleteCategoryByID(id int) error {
	err := cs.RepoCategory.DeleteCategoryByID(id)
	return err
}
