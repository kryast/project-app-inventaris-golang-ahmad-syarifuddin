package handler

import "net/http"

func (ih *ItemHandler) CMSAllItemHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil data item
	items, err := ih.serviceItems.GetAllItem()
	if err != nil {
		http.Error(w, "Gagal mengambil data item", http.StatusInternalServerError)
		return
	}

	// Render template dengan data items
	templates.ExecuteTemplate(w, "item-list-view", items)

}
