package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

func (ih *ItemHandler) GetItemInvestmentByIDHandler(w http.ResponseWriter, r *http.Request) {

	itemID := library.GetID(w, r)

	itemInvestment, err := ih.serviceItems.GetItemInvestmentByID(itemID)
	if err != nil {
		library.Response500(w, err.Error())
		return
	}

	library.Response200(w, itemInvestment)
}
