package main

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/database"
	"project-app-inventaris-golang-ahmad-syarifuddin/handler"
	"project-app-inventaris-golang-ahmad-syarifuddin/library"
	"project-app-inventaris-golang-ahmad-syarifuddin/middleware"
	"project-app-inventaris-golang-ahmad-syarifuddin/repository"
	"project-app-inventaris-golang-ahmad-syarifuddin/service"

	"github.com/go-chi/chi"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoItem := repository.NewItemRepository(db)
	serviceItem := service.NewItemService(repoItem)
	itemHandler := handler.NewItemHandler(serviceItem)

	repoCategory := repository.NewCategoryRepository(db)
	serviceCategory := service.NewCategoryService(repoCategory)
	categoryHandler := handler.NewCategoryHandler(serviceCategory)

	repoInvestment := repository.NewInvestmentRepository(db)
	serviceInvestment := service.NewInvestmentService(repoInvestment)
	investmentHandler := handler.NewInvestmentHandler(serviceInvestment)

	repoAdmin := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(repoAdmin)
	adminHandler := handler.NewAdminHandler(adminService)

	r := chi.NewRouter()

	r.Use(library.MethodForm)

	// Rute Postman
	r.Get("/api/items", itemHandler.GetAllItemHandler)
	// r.Get("/api/items", itemHandler.GetItemsWithFiltersHandler)
	r.Get("/api/items/{id}", itemHandler.GetItemByIdHandler)
	r.Post("/api/items", itemHandler.CreateItemHandler)
	r.Put("/api/items/{id}", itemHandler.UpdateItemHandler)
	r.Delete("/api/items/{id}", itemHandler.DeleteItemHandler)

	r.Get("/api/items/replacement-needed", itemHandler.GetItemReplacementHandler)

	r.Get("/api/items/investment", itemHandler.GetAllInvestmentHandler)
	r.Get("/api/items/investment/{id}", itemHandler.GetItemInvestmentByIDHandler)

	r.Post("/api/categories", categoryHandler.CreateCategoryHandler)
	r.Get("/api/categories", categoryHandler.GetAllCategoryHandler)
	r.Get("/api/categories/{id}", categoryHandler.GetCategoryByIdHandler)
	r.Put("/api/categories/{id}", categoryHandler.UpdateCategoryHandler)
	r.Delete("/api/categories/{id}", categoryHandler.DeleteCategoryHandler)

	// Rute CMS

	r.Get("/login", handler.FormLogin)
	r.Post("/login", adminHandler.LoginHandler)

	r.Get("/register-view", handler.FormRegister)
	r.Post("/register", adminHandler.CreateAdminHandler)

	r.With(middleware.CheckLoginMiddleware).Group(func(r chi.Router) {

		r.Get("/", itemHandler.Home)
		r.Get("/create-item", categoryHandler.CMSCreateItemPageHandler)
		r.Post("/create-item", itemHandler.CMSCreateItemHandler)
		r.Get("/all-item", itemHandler.CMSAllItemHandler)
		r.Get("/update-item/{id}", itemHandler.CMSUpdateItemPageHandler)
		r.Put("/update-item/{id}", itemHandler.CMSUpdateItemHandler)
		r.Delete("/delete-item/{id}", itemHandler.CMSDeleteItemHandler)

		r.Get("/investment", investmentHandler.CMSGetItemInvestmentHandler)
		r.Get("/logout-view", handler.Logout)
		r.Get("/logout", adminHandler.LogoutHandler)
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

}
