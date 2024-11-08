package service

import "project-app-inventaris-golang-ahmad-syarifuddin/model"

func (cs *CategoryService) UpdateCategory(category model.Category) error {
	return cs.RepoCategory.UpdateCategoryByID(category)
}
