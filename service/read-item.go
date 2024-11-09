package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (is *ItemService) GetAllItem() ([]model.Item, error) {

	items, err := is.RepoItem.GetAllItem()
	if err != nil {
		return nil, fmt.Errorf("error at service")
	}

	return items, nil

}
