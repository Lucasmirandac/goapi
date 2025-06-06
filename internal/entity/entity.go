package entity

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func NewProduct(name string, price float64, description string, categoryID string, imageURL string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Price:       price,
		Description: description,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
