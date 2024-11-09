package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"strconv"
	"time"
)

func (ih *ItemHandler) CMSCreateItemHandler(w http.ResponseWriter, r *http.Request) {
	// Pastikan request adalah POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Ambil data dari form
	name := r.FormValue("item")               // Nama item
	category := r.FormValue("category")       // ID kategori yang dipilih
	categoryID, err := strconv.Atoi(category) // Konversi kategori ke ID
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Ambil harga dan tanggal pembelian
	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr) // Konversi harga menjadi integer
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	purchaseDateStr := r.FormValue("purchase_date")
	layout := "2006-01-02"
	purchaseDate, err := time.Parse(layout, purchaseDateStr) // Parse tanggal pembelian
	if err != nil {
		http.Error(w, "Invalid purchase date format. Expected format: YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// Hitung total penggunaan hari
	totalUsageDays := int(time.Since(purchaseDate).Hours() / 24)

	// Tentukan apakah perlu penggantian
	var replacementRequired bool
	if totalUsageDays > 100 {
		replacementRequired = true
	} else {
		replacementRequired = false
	}

	// Proses upload file gambar
	photoPath, err := library.UploadFile(r, "photo_url", "./uploads", "jpg")
	if err != nil {
		http.Error(w, "Error uploading photo file", http.StatusInternalServerError)
		return
	}

	// Buat item baru dengan data yang diambil
	item := model.Item{
		Name:                name,
		CategoryID:          categoryID, // ID kategori yang dipilih
		PhotoURL:            photoPath,
		Price:               price,
		PurchaseDate:        purchaseDate,
		TotalUsageDay:       totalUsageDays,
		ReplacementRequired: replacementRequired,
	}

	// Panggil service untuk menyimpan item baru
	err = ih.serviceItems.CreateItem(item)
	if err != nil {
		http.Error(w, "Error saving item to the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect ke halaman utama atau ke halaman lain setelah sukses
	http.Redirect(w, r, "/all-item", http.StatusSeeOther)
}
