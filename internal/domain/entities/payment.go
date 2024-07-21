package entities

import "github.com/google/uuid"

type PaymentStatus struct {
	ID     uuid.UUID
	Status int
}
