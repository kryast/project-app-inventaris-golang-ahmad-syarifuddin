package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	_, err := ih.serviceItems.GetItemById(id)
	if err != nil {
		library.Response500(w, fmt.Sprintf("id: %d not found", id))
		return
	}

	err = ih.serviceItems.DeleteItemByID(id)
	if err != nil {
		library.Response500(w, fmt.Sprintf("Error deleting item with id: %d", id))
		return
	}

	library.ResponseDelete200(w, "Item Berhasil Dihapus")
}
