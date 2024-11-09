package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) GetItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	item, err := ih.serviceItems.GetItemById(id)
	if err != nil {
		library.Response404(w, fmt.Sprintf("item with id: %d not found", id))
		return
	}

	library.Response200(w, item)
}
