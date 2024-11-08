package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/repository"
)

type ItemService struct {
	RepoItem repository.ItemRepositoryDB
}

func NewItemService(repo repository.ItemRepositoryDB) ItemService {
	return ItemService{RepoItem: repo}
}

func (is *ItemService) CreateItem(item model.Item) error {
	err := is.RepoItem.CreateDataItem(item)
	if err != nil {
		return fmt.Errorf("error while creating item in service: %v", err)
	}
	return nil
}
