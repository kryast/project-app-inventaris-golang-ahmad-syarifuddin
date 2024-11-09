package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) CMSUpdateItemPageHandler(w http.ResponseWriter, r *http.Request) {

	id := library.GetID(w, r)

	item, err := ih.serviceItems.GetItemById(id)
	if err != nil {
		http.Error(w, "Error fetching item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "update-item-view", item)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
