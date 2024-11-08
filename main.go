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

	r := chi.NewRouter()

	// Rute untuk User
	r.Post("/create", itemHandler.CreateItemHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

}
