package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ch *CategoryHandler) GetAllCategoryHandler(w http.ResponseWriter, r *http.Request) {

	categories, err := ch.serviceCategories.GetAllCategory()
	if err != nil {
		library.Response404(w, "Kategori tidak ditemukan")
		return
	}

	library.Response200(w, categories)
}
