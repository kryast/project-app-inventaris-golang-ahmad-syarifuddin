package handler

import (
	"encoding/json"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (ih *ItemHandler) UpdateItemHandler(w http.ResponseWriter, r *http.Request) {

	id := library.GetID(w, r)

	oldItem, err := ih.serviceItems.GetItemById(id)
	if err != nil {
		library.Response400(w, "Item not found")
		return
	}

	var item model.Item
	err = json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		library.Response400(w, "Invalid request data")
		return
	}
	if item.Name == "" {
		item.Name = oldItem.Name
	}
	if item.CategoryID == 0 {
		item.CategoryID = oldItem.CategoryID
	}
	if item.PhotoURL == "" {
		item.PhotoURL = oldItem.PhotoURL
	}
	if item.Price == 0 {
		item.Price = oldItem.Price
	}
	if item.PurchaseDate.IsZero() {
		item.PurchaseDate = oldItem.PurchaseDate
	}
	if item.TotalUsageDay == 0 {
		item.TotalUsageDay = oldItem.TotalUsageDay
	}
	if item.ReplacementRequired == nil {
		item.ReplacementRequired = oldItem.ReplacementRequired
	}

	item.ID = id
	err = ih.serviceItems.UpdateItem(item)
	if err != nil {
		library.Response400(w, "Failed to update category")
		return
	}
	library.Response200(w, item)
}
