package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (is *ItemService) GetItemById(id int) (*model.Item, error) {
	item, err := is.RepoItem.GetItemById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user details: %v", err)
	}
	return item, nil
}
