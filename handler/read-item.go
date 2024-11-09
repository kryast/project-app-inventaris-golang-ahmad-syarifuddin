package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) GetAllItemHandler(w http.ResponseWriter, r *http.Request) {

	items, err := ih.serviceItems.GetAllItem()
	if err != nil {
		library.Response404(w, "Item tidak ditemukan")
		return
	}

	library.Response200(w, items)
}
