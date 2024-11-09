package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (ih *ItemHandler) CMSUpdateItemHandler(w http.ResponseWriter, r *http.Request) {

	itemIDStr := chi.URLParam(r, "id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil || itemID <= 0 {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	oldItem, err := ih.serviceItems.GetItemById(itemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Item with ID %d not found", itemID), http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		name = oldItem.Name
	}

	categoryIDStr := r.FormValue("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil || categoryID <= 0 {
		categoryID = oldItem.CategoryID
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil || price < 0 {
		price = oldItem.Price
	}

	purchaseDateStr := r.FormValue("purchase_date")
	var purchaseDate time.Time
	if purchaseDateStr != "" {
		purchaseDate, err = time.Parse("2006-01-02", purchaseDateStr)
		if err != nil {
			http.Error(w, "Invalid purchase date format", http.StatusBadRequest)
			return
		}
	} else {
		purchaseDate = oldItem.PurchaseDate
	}

	photoURL := oldItem.PhotoURL
	if _, _, err := r.FormFile("photo_url"); err == nil {
		photoURL, err = library.UploadFile(r, "photo_url", "./uploads/items", "jpg")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error uploading photo: %v", err), http.StatusInternalServerError)
			return
		}
	}

	item := model.Item{
		ID:           itemID,
		Name:         name,
		CategoryID:   categoryID,
		Price:        price,
		PurchaseDate: purchaseDate,
		PhotoURL:     photoURL,
	}

	err = ih.serviceItems.UpdateItem(item)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating item: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/all-item", http.StatusSeeOther)
}
