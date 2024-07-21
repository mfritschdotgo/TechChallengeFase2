package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                uuid.UUID   `bson:"_id" json:"id"`
	Order             int         `bson:"order" json:"order"`
	Client            string      `bson:"client" json:"client"`
	Items             []OrderItem `bson:"items" json:"items"`
	Total             float64     `bson:"total" json:"total"`
	Status            int         `bson:"status" json:"status"`
	StatusDescription string      `bson:"status_description" json:"status_description"`
	CreatedAt         time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time   `bson:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	ProductID   string  `bson:"product_id" json:"product_id"`
	ProductName string  `bson:"product_name" json:"product_name"`
	Quantity    int64   `bson:"quantity" json:"quantity"`
	Price       float64 `bson:"price" json:"price"`
}

type OrderStatus struct {
	Status            int    `json:"status"`
	StatusDescription string ` json:"status_description"`
}

func NewOrder(client string, items []OrderItem, status int, total float64, statusDescription string) (*Order, error) {
	now := time.Now()

	order := &Order{
		ID:                uuid.New(),
		Client:            client,
		Items:             items,
		Total:             total,
		Status:            status,
		StatusDescription: statusDescription,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return order, nil
}

func SetStatus(status int) (*OrderStatus, error) {
	var statusDescription string

	if status == 1 {
		statusDescription = "received"
	} else if status == 2 {
		statusDescription = "preparing"
	} else if status == 3 {
		statusDescription = "ready"
	} else if status == 4 {
		statusDescription = "finished"
	} else {
		return nil, fmt.Errorf("invalid status")
	}

	OrderStatus := &OrderStatus{
		Status:            status,
		StatusDescription: statusDescription,
	}

	return OrderStatus, nil
}
