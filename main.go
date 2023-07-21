package main

import (
	_ "github.com/lib/pq"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	/*
		Without DI
		db := app.NewDB()
		validate := validator.New()
		categoryRepository := repository.NewCategoryRepository()
		categoryService := service.NewCategoryService(categoryRepository, db, validate)
		categoryController := controller.NewCategoryController(categoryService)
		router := app.NewRouter(categoryController)
		authMiddleware := middleware.NewAuthMiddleware(router)
		server := NewServer(authMiddleware)
	*/

	// with DI
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
