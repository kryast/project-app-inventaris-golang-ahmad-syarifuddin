package handler

import (
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/service"
	"strconv"
	"time"
)

type AdminHandler struct {
	serviceAdmin service.AdminService
}

func NewAdminHandler(service service.AdminService) *AdminHandler {
	return &AdminHandler{serviceAdmin: service}
}

func (ah *AdminHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	admin, err := ah.serviceAdmin.ValidateAdmin(username, password)
	if err != nil {
		library.Response400(w, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:    "admin-session",
		Value:   strconv.Itoa(admin.ID),
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
