package handler

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) GetAllInvestmentHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil summary total investasi dan depresiasi dari service
	investmentSummary, err := ih.serviceItems.GetAllInvestment()
	if err != nil {
		library.Response500(w, fmt.Sprintf("Error fetching investment summary: %v", err))
		return
	}

	library.Response200(w, investmentSummary)
}
