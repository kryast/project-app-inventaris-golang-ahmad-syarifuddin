package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (ah *AdminHandler) CreateAdminHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if name == "" || username == "" || password == "" {
		library.Response400(w, "All fields are required")
		return
	}

	admin := model.Admin{
		Name:     name,
		Username: username,
		Password: password,
	}

	err := ah.serviceAdmin.CreateAdmin(admin)
	if err != nil {
		library.Response400(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
