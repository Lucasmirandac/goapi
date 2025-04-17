package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Lucasmirandac/go_api_ecommerce/internal/database"
	"github.com/Lucasmirandac/go_api_ecommerce/internal/service"
	"github.com/Lucasmirandac/go_api_ecommerce/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/catalog")
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)
	webCategoryHandler := webserver.NewWebCategoryHandler(*categoryService)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	webProductHandler := webserver.NewWebProductHandler(*productService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Route("/products", func(r chi.Router) {
		r.Get("/", webProductHandler.GetProducts)
		r.Get("/{id}", webProductHandler.GetProduct)
		r.Get("/category/{categoryID}", webProductHandler.GetProductsByCategoryID)
		r.Post("/", webProductHandler.CreateProduct)
		r.Put("/{id}", webProductHandler.UpdateProduct)
		r.Delete("/{id}", webProductHandler.DeleteProduct)
	})

	router.Route("/categories", func(r chi.Router) {
		r.Get("/", webCategoryHandler.GetCategories)
		r.Get("/{id}", webCategoryHandler.GetCategory)
		r.Post("/", webCategoryHandler.CreateCategory)
	})

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
