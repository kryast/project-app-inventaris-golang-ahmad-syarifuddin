package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"project-app-inventaris-golang-ahmad-syarifuddin/library"
// 	"project-app-inventaris-golang-ahmad-syarifuddin/model"
// 	"strconv"
// )

// func (ih *ItemHandler) GetItemsWithFiltersHandler(w http.ResponseWriter, r *http.Request) {
// 	defaultLimit := 2
// 	page := 1
// 	category := r.URL.Query().Get("category")
// 	search := r.URL.Query().Get("search")
// 	totalUsageDays := r.URL.Query().Get("total-usage-days")
// 	pageStr := r.URL.Query().Get("page")
// 	limitStr := r.URL.Query().Get("limit")

// 	if pageStr != "" {
// 		page, _ = strconv.Atoi(pageStr)
// 	}
// 	if limitStr != "" {
// 		defaultLimit, _ = strconv.Atoi(limitStr)
// 	}

// 	filter := model.ItemFilter{
// 		Category:       category,
// 		Search:         search,
// 		TotalUsageDays: totalUsageDays,
// 		Page:           page,
// 		Limit:          defaultLimit,
// 	}

// 	items, totalItems, err := ih.serviceItems.GetFilteredItems(filter)
// 	if err != nil {
// 		library.Response500(w, fmt.Sprintf("Error fetching items: %v", err))
// 		return
// 	}

// 	totalPages := (totalItems + defaultLimit - 1) / defaultLimit

// 	response := struct {
// 		Success    bool         `json:"success"`
// 		Page       int          `json:"page"`
// 		Limit      int          `json:"limit"`
// 		TotalItems int          `json:"total_items"`
// 		TotalPages int          `json:"total_pages"`
// 		Data       []model.Item `json:"data"`
// 	}{
// 		Success:    true,
// 		Page:       page,
// 		Limit:      defaultLimit,
// 		TotalItems: totalItems,
// 		TotalPages: totalPages,
// 		Data:       items,
// 	}

// 	successResponse, _ := json.MarshalIndent(response, "", " ")
// 	jsonResponse, _ := w.Write(successResponse)
// 	json.NewEncoder(w).Encode(jsonResponse)
// }
