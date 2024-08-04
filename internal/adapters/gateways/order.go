package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type OrderGateway struct {
	repo interfaces.OrderRepository
}

func NewOrderGateway(repo interfaces.OrderRepository) interfaces.OrderGateway {
	return &OrderGateway{
		repo: repo,
	}
}

func (g *OrderGateway) CreateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	return g.repo.CreateOrder(ctx, order)
}

func (g *OrderGateway) GetOrders(ctx context.Context, page, limit int) ([]entities.Order, error) {
	return g.repo.GetOrders(ctx, page, limit)
}

func (g *OrderGateway) GetOrderByID(ctx context.Context, id string) (*entities.Order, error) {
	return g.repo.GetOrderByID(ctx, id)
}

func (g *OrderGateway) SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error {
	return g.repo.SetStatus(ctx, id, status, description)
}
