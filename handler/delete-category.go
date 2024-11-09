package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ch *CategoryHandler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	_, err := ch.serviceCategories.GetCategoryById(id)
	if err != nil {
		library.Response500(w, fmt.Sprintf("id: %d not found", id))
		return
	}

	err = ch.serviceCategories.DeleteCategoryByID(id)
	if err != nil {
		library.Response500(w, fmt.Sprintf("Error deleting category with id: %d", id))
		return
	}

	library.ResponseDelete200(w, "Kategori Berhasil Dihapus")
}
