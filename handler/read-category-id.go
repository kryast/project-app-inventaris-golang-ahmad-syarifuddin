package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ch *CategoryHandler) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	category, err := ch.serviceCategories.GetCategoryById(id)
	if err != nil {
		library.Response404(w, fmt.Sprintf("category: %d not found", id))
		return
	}

	library.Response200(w, category)
}
