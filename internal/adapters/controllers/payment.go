package controllers

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type PaymentController struct {
	usecase   interfaces.PaymentUseCase
	presenter interfaces.PaymentPresenter
}

func NewPaymentController(u interfaces.PaymentUseCase, p interfaces.PaymentPresenter) *PaymentController {
	return &PaymentController{
		usecase:   u,
		presenter: p,
	}
}

func (c *PaymentController) SetPaymentStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error) {
	return c.usecase.SetPaymentStatus(ctx, id, status)
}

func (c *PaymentController) GenerateQRCode(ctx context.Context, id string) ([]byte, error) {
	return c.usecase.GenerateQRCode(ctx, id)
}
