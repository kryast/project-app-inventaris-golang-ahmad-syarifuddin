package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (is *ItemService) GetItemsForReplacement() ([]model.Item, error) {
	items, err := is.RepoItem.GetItemsForReplacement()
	if err != nil {
		fmt.Printf("Error in GetItemsForReplacement: %v\n", err)
		return nil, err
	}

	return items, nil
}
