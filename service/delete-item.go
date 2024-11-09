package service

func (is *ItemService) DeleteItemByID(id int) error {
	err := is.RepoItem.DeleteItemByID(id)
	return err
}
