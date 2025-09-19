package main

import (
	"api/order/configs"
	"api/order/internal/auth"
	"api/order/internal/product"
	"api/order/internal/user"
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
	productRepo := product.NewProductRepository(db)
	userRepo := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthService(userRepo)

	// Handlers
	product.NewProductHandler(router, product.ProductHandlerDeps{
		Config:            config,
		ProductRepository: productRepo,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
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
}
