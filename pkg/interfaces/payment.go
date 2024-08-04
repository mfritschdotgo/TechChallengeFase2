package interfaces

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
)

type PaymentRepository interface {
	UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error
}

type PaymentGateway interface {
	UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error
}

type PaymentUseCase interface {
	SetPaymentStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error)
	GenerateQRCode(ctx context.Context, id string) ([]byte, error)
}

type PaymentPresenter interface {
	JSON(w http.ResponseWriter, data interface{}, statusCode int)
	Error(w http.ResponseWriter, message string, statusCode int)
}
