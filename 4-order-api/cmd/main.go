package main

import (
	"api/order/configs"
	"api/order/internal/order"
	"api/order/migrations"
	"api/order/pkg/db"
	"api/order/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	// Initialize the application
	migrations.Migrate()
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()
	// Repositories
	productRepo := order.NewProductRepository(db)
	// Handlers
	order.NewProductHandler(router, order.ProductHandlerDeps{
		ProductRepository: productRepo,
	})

	stack := middleware.Chain(
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is running on port 8081...")
	server.ListenAndServe()

	// Further initialization logic can go here
}
