package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (ch *CategoryHandler) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	id := library.GetID(w, r)

	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		library.Response400(w, "Invalid request data")
		return
	}

	category.ID = id
	err = ch.serviceCategories.UpdateCategory(category)
	if err != nil {
		library.Response400(w, "Failed to update category")
		return
	}
	library.Response200(w, category)
}
