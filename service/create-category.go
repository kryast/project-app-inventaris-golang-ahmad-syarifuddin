package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/repository"
)

type CategoryService struct {
	RepoCategory repository.CategoryRepositoryDB
}

func NewCategoryService(repo repository.CategoryRepositoryDB) CategoryService {
	return CategoryService{RepoCategory: repo}
}

func (cs *CategoryService) CreateCategory(category model.Category) error {
	err := cs.RepoCategory.CreateDataCategory(category)
	if err != nil {
		return fmt.Errorf("error at service")
	}
	return nil
}
