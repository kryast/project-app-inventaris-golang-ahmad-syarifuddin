package handler

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "logout-view", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home-view", nil)
}
func CreateItemView(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create-item-view", nil)
}
