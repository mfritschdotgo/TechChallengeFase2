package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, product *entities.Order) (*entities.Order, error)
	GetOrderByID(ctx context.Context, id string) (*entities.Order, error)
	GetOrders(ctx context.Context, page, pageSize int) ([]entities.Order, error)
	SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error
}
