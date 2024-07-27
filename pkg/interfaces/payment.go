package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type PaymentStatusRepository interface {
	UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error
}
