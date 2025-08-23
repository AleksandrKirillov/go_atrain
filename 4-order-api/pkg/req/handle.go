package req

import (
	"api/order/pkg/resp"
	"fmt"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	fmt.Println("Handling request body...")
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	fmt.Println("Decoded body:", body)

	if err := IsValid(body); err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}
