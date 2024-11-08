package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/service"
)

type CategoryHandler struct {
	serviceCategories service.CategoryService
}

func NewCategoryHandler(cs service.CategoryService) CategoryHandler {
	return CategoryHandler{serviceCategories: cs}
}

func (ch *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		library.Response400(w, "Invalid request body")
		return
	}

	err = ch.serviceCategories.CreateCategory(category)
	if err != nil {
		library.Response500(w, "Failed to create category")
		return
	}
	library.Response201(w, "Category created successfully", category)
}
