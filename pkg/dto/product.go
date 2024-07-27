package dto

import "github.com/google/uuid"

type CreateProductRequest struct {
	CategoryId  uuid.UUID `json:"category_id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}
