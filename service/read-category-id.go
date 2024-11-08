package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (cs *CategoryService) GetCategoryById(id int) (*model.Category, error) {
	category, err := cs.RepoCategory.GetCategoryById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user details: %v", err)
	}
	return category, nil
}
