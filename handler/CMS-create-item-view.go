package handler

import (
	"net/http"
)

func (ih *CategoryHandler) CMSCreateItemPageHandler(w http.ResponseWriter, r *http.Request) {

	categories, err := ih.serviceCategories.GetAllCategory()
	if err != nil {
		http.Error(w, "Error fetching categories: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "create-item-view", categories)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
