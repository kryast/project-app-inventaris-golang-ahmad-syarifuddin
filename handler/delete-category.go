package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ch *CategoryHandler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	err := ch.serviceCategories.DeleteCategoryByID(id)
	if err != nil {
		library.Response500(w, fmt.Sprintf("Error deleting category: %d", id))
		return
	}

	library.ResponseDelete200(w, "Kategori Berhasil Dihapus")
}
