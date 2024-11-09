package main

import (
	"fmt"
	"net/http"
	"project-app-inventaris-golang-ahmad-syarifuddin/database"
	"project-app-inventaris-golang-ahmad-syarifuddin/handler"
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

	r := chi.NewRouter()

	// Rute untuk User
	r.Get("/api/items", itemHandler.GetAllItemHandler)
	r.Get("/api/items/{id}", itemHandler.GetItemByIdHandler)
	r.Post("/api/items", itemHandler.CreateItemHandler)
	r.Put("/api/items/{id}", itemHandler.UpdateItemHandler)

	r.Post("/api/categories", categoryHandler.CreateCategoryHandler)
	r.Get("/api/categories", categoryHandler.GetAllCategoryHandler)
	r.Get("/api/categories/{id}", categoryHandler.GetCategoryByIdHandler)
	r.Put("/api/categories/{id}", categoryHandler.UpdateCategoryHandler)
	r.Delete("/api/categories/{id}", categoryHandler.DeleteCategoryHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

}
