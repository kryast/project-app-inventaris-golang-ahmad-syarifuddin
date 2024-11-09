package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/service"
)

type InvestmentHandler struct {
	serviceInvestment service.InvestmentService
}

func NewInvestmentHandler(cs service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{serviceInvestment: cs}
}

func (ih *InvestmentHandler) CMSGetItemInvestmentHandler(w http.ResponseWriter, r *http.Request) {

	itemInvestment, err := ih.serviceInvestment.GetAllInvestment()
	if err != nil {
		library.Response500(w, err.Error())
		return
	}

	templates.ExecuteTemplate(w, "investment-item-view", itemInvestment)

}
