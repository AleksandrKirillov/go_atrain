package order

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}
type ProductUpdateRequest struct {
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}
