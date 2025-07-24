package main

import (
	"api/validation/configs"
	auth "api/validation/internal"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
