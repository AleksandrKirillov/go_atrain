package order

import (
	"api/order/internal/product"
	"api/order/internal/user"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User     user.User         // belongs to User
	Products []product.Product `gorm:"many2many:order_products;"` // N:N
}

func NewOrder(name string, description string, images pq.StringArray) *Order {
	return &Order{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
