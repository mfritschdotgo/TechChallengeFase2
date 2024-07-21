package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" bson:"_id"`
	CategoryId  uuid.UUID `json:"category_id" bson:"category_id"`
	Name        string    `json:"name" bson:"name"`
	Price       float64   `json:"price" bson:"price"`
	Description string    `json:"description" bson:"description"`
	Image       string    `json:"image" bson:"image"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func NewProduct(name string, price float64, categoryId uuid.UUID, description string, image string) (*Product, error) {
	now := time.Now()

	product := &Product{
		ID:          uuid.New(),
		CategoryId:  categoryId,
		Name:        name,
		Price:       price,
		Description: description,
		Image:       image,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return product, nil
}
