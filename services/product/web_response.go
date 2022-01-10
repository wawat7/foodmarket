package product

import (
	"encoding/json"
	"time"
)

type FormatProduct struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ingredient  []string  `json:"ingredient"`
	Price       int       `json:"price"`
	Rate        float32   `json:"rate"`
	Type        string    `json:"type"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ProductFormat(product Product) FormatProduct {
	var ingredients []string
	json.Unmarshal([]byte(product.Ingredient), &ingredients)

	return FormatProduct{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Ingredient:  ingredients,
		Price:       product.Price,
		Rate:        product.Rate,
		Type:        product.Type,
		Image:       product.Image,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ProductsFormat(products []Product) []FormatProduct {
	var formats []FormatProduct

	for _, product := range products {
		format := ProductFormat(product)
		formats = append(formats, format)
	}

	return formats
}
