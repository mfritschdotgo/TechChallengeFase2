package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type PaymentGateway struct {
	repo interfaces.PaymentRepository
}

func NewPaymentGateway(repo interfaces.PaymentRepository) interfaces.PaymentGateway {
	return &PaymentGateway{
		repo: repo,
	}
}

func (g *PaymentGateway) UpdatePayment(ctx context.Context, id uuid.UUID, status int, description string) error {
	return g.repo.UpdatePayment(ctx, id, status, description)
}
