package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) CMSDeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := library.GetID(w, r)

	err := ih.serviceItems.DeleteItemByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting book: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/all-item", http.StatusSeeOther)
}
