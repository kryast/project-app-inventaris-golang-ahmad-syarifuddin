package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) GetItemReplacementHandler(w http.ResponseWriter, r *http.Request) {
	items, err := ih.serviceItems.GetItemsForReplacement()
	if err != nil {
		library.Response500(w, "Error fetching items for replacement")
		return
	}

	if len(items) == 0 {
		library.Response404(w, "No items found that require replacement")
		return
	}

	library.Response200(w, items)
}
