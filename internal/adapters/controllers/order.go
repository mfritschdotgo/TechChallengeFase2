package controllers

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type OrderController struct {
	usecase   interfaces.OrderUseCase
	presenter interfaces.OrderPresenter
}

func NewOrderController(u interfaces.OrderUseCase, p interfaces.OrderPresenter) *OrderController {
	return &OrderController{
		usecase:   u,
		presenter: p,
	}
}

func (c *OrderController) CreateOrder(ctx context.Context, orderDto dto.CreateOrderRequest) (*entities.Order, error) {
	return c.usecase.CreateOrder(ctx, orderDto)
}

func (c *OrderController) GetOrderByID(ctx context.Context, id string) (*entities.Order, error) {
	return c.usecase.GetOrderByID(ctx, id)
}

func (c *OrderController) GetOrders(ctx context.Context, page, size int) ([]entities.Order, error) {
	return c.usecase.GetOrders(ctx, page, size)
}

func (c *OrderController) SetOrderStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error) {
	return c.usecase.SetOrderStatus(ctx, id, status)
}
