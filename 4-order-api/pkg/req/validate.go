package req

import (
	"fmt"

	"github.com/go-playground/validator"
)

func IsValid[T any](payload T) error {
	validate := validator.New()
	fmt.Println("Validating payload:", payload)
	if err := validate.Struct(payload); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	fmt.Println("Payload is valid")

	return nil
}
