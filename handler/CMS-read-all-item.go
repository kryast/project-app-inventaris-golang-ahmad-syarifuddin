package handler

import "net/http"

func (ih *ItemHandler) CMSAllItemHandler(w http.ResponseWriter, r *http.Request) {
	
	items, err := ih.serviceItems.GetAllItem()
	if err != nil {
		http.Error(w, "Gagal mengambil data item", http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "item-list-view", items)

}
