package handler

import (
	"html/template"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "logout-view", nil)
}

func (ih *ItemHandler) Home(w http.ResponseWriter, r *http.Request) {
	totalInvestment, err := ih.serviceItems.GetTotalInvestment()
	if err != nil {
		library.Response400(w, err.Error())
	}
	templates.ExecuteTemplate(w, "home-view", totalInvestment)
}
func CreateItemView(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-item-view", nil)
}
