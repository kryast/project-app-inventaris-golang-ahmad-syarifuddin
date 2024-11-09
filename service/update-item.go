package service

import "project-app-inventaris-golang-ahmad-syarifuddin/model"

func (is *ItemService) UpdateItem(item model.Item) error {
	return is.RepoItem.UpdateItemByID(item)
}
