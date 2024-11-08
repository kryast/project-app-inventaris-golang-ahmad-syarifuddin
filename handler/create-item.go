package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/service"
	"strconv"
	"time"
)

type ItemHandler struct {
	serviceItems service.ItemService
}

func NewItemHandler(bs service.ItemService) ItemHandler {
	return ItemHandler{serviceItems: bs}
}

func (ih *ItemHandler) CreateItemHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	category := r.FormValue("category")
	categoryID, _ := strconv.Atoi(category)

	photoPath, err := library.UploadFile(r, "photo_url", "./uploads", "jpg")
	if err != nil {
		library.Response500(w, "Error uploading photo_url file")
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		library.Response400(w, "Invalid price value")
		return
	}

	purchaseDateStr := r.FormValue("purchase_date")

	layout := "2006-01-02"
	purchaseDate, err := time.Parse(layout, purchaseDateStr)
	if err != nil {
		library.Response400(w, "Invalid purchase date format. Expected format: YYYY-MM-DD")
		return
	}
	totalUsageDays := int(time.Since(purchaseDate).Hours() / 24)

	item := model.Item{
		Name:          name,
		Category:      categoryID,
		PhotoURL:      photoPath,
		Price:         price,
		PurchaseDate:  purchaseDate,
		TotalUsageDay: totalUsageDays,
	}

	err = ih.serviceItems.CreateItem(item)
	if err != nil {
		library.Response500(w, err.Error())
		return
	}

	library.Response201(w, "Item created successfully", item)
}
