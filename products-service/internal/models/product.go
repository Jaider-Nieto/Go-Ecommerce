package models

import (
	"github.com/jaider-nieto/ecommerce-go/products-service/pkg/utils"
)

// swagger:model
type Product struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Category    string `json:"category" bson:"category"`
	Price       uint   `json:"price" bson:"price"`
	Stock       uint   `json:"stock" bson:"stock"`
	Image       string `json:"image" bson:"image"`
	Rating      []uint `json:"rating" bson:"rating"`
}

type CreateProduct struct {
	Title       string `json:"title" bson:"title" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
	Category    string `json:"category" bson:"category" validate:"required"`
	Price       uint   `json:"price" bson:"price" validate:"required"`
	Stock       uint   `json:"stock" bson:"stock" validate:"required"`
	Image       string `bson:"stock"`
}

func (p *CreateProduct) IsValidCategory() bool {
	return utils.IsValidCategory(p.Category)
}
