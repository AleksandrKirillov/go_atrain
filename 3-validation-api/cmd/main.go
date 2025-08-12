package main

import (
	"api/validation/configs"
	auth "api/validation/internal"
	"api/validation/storage"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	storage := storage.NewStorage("storage.json")
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:  conf,
		Storage: storage,
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
