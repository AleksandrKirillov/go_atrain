package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/random", random)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}

func random(w http.ResponseWriter, r *http.Request) {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	randomNum := random.Intn(6) + 1 // Random number between 1 and 6
	numb := []byte(fmt.Sprintf("%d", randomNum))
	w.Write(numb)
}
