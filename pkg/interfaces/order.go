package interfaces

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	GetOrders(ctx context.Context, page, limit int) ([]entities.Order, error)
	GetOrderByID(ctx context.Context, id string) (*entities.Order, error)
	SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error
}

type OrderGateway interface {
	CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	GetOrders(ctx context.Context, page, limit int) ([]entities.Order, error)
	GetOrderByID(ctx context.Context, id string) (*entities.Order, error)
	SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error
}

type OrderUseCase interface {
	CreateOrder(ctx context.Context, dto dto.CreateOrderRequest) (*entities.Order, error)
	GetOrderByID(ctx context.Context, id string) (*entities.Order, error)
	GetOrders(ctx context.Context, page, size int) ([]entities.Order, error)
	SetOrderStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error)
}

type OrderPresenter interface {
	JSON(w http.ResponseWriter, data interface{}, statusCode int)
	Error(w http.ResponseWriter, message string, statusCode int)
}
