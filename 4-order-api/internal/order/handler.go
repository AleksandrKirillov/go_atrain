package order

import (
	"api/order/configs"
	"api/order/pkg/middleware"
	"api/order/pkg/req"
	"api/order/pkg/resp"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	*configs.Config
	ProductRepository *ProductRepository
}
type ProductHandler struct {
	*configs.Config
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		Config:            deps.Config,
		ProductRepository: deps.ProductRepository,
	}

	router.Handle("POST /product", middleware.Auth(handler.Create(),
		middleware.AuthMiddleware{Secret: handler.Config.Auth.Secret}))
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/{id}", handler.Get())

}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextSessionKey).(string)
		if !ok || phone == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("Creating product...")
		body, err := req.HandleBody[ProductCreateRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Product body:", body)
		product := NewProduct(body.Name, body.Description, body.Images)
		fmt.Println("Product created:", product)
		createdProduct, err := handler.ProductRepository.Create(product)
		fmt.Println("Created product:", createdProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Json(w, createdProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.Get(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Json(w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Description: body.Description,
			Images:      body.Images,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Json(w, product, http.StatusOK)
	}
}
func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			if err.Error() == "Запись не найдена" {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		resp.Json(w, nil, http.StatusOK)
	}
}
