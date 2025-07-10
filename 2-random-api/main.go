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
	numb := []byte(fmt.Sprintf("%d", random.Intn(10)))
	w.Write(numb)
}
