package main

import (
	"api/order/configs"
	"api/order/migrations"
	"api/order/pkg/db"
)

func main() {
	// Initialize the application
	migrations.Migrate()
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	// Further initialization logic can go here
}
