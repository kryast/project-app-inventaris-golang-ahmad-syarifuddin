package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"strconv"
	"time"
)

func (ih *ItemHandler) CMSCreateItemHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("item")
	category := r.FormValue("category")
	categoryID, err := strconv.Atoi(category)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	purchaseDateStr := r.FormValue("purchase_date")
	layout := "2006-01-02"
	purchaseDate, err := time.Parse(layout, purchaseDateStr)
	if err != nil {
		http.Error(w, "Invalid purchase date format. Expected format: YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	totalUsageDays := int(time.Since(purchaseDate).Hours() / 24)

	var replacementRequired bool
	if totalUsageDays > 100 {
		replacementRequired = true
	} else {
		replacementRequired = false
	}

	photoPath, err := library.UploadFile(r, "photo_url", "./uploads", "jpg")
	if err != nil {
		http.Error(w, "Error uploading photo file", http.StatusInternalServerError)
		return
	}

	item := model.Item{
		Name:                name,
		CategoryID:          categoryID,
		PhotoURL:            photoPath,
		Price:               price,
		PurchaseDate:        purchaseDate,
		TotalUsageDay:       totalUsageDays,
		ReplacementRequired: replacementRequired,
	}

	err = ih.serviceItems.CreateItem(item)
	if err != nil {
		http.Error(w, "Error saving item to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/all-item", http.StatusSeeOther)
}
