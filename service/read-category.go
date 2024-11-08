package service

import "project-app-inventaris-golang-ahmad-syarifuddin/model"

func (cs *CategoryService) GetAllCategory() (*[]model.Category, error) {

	categories, err := cs.RepoCategory.GetAllCategory()

	if err != nil {
		return nil, err
	}

	return categories, nil

}
