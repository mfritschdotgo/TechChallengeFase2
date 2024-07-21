package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/interfaces"
	"github.com/skip2/go-qrcode"
)

type Payment struct {
	paymentRepo   interfaces.PaymentStatusRepository
	orderUseCases *Order
}

func NewPaymentStatusUsecase(repo interfaces.PaymentStatusRepository, order *Order) *Payment {
	return &Payment{
		paymentRepo:   repo,
		orderUseCases: order,
	}
}

func (s *Payment) SetPaymentStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	if status > 1 || status < 0 {
		return nil, fmt.Errorf("invalid status")
	}

	order, err := s.orderUseCases.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	if order.Status != 0 {
		return nil, fmt.Errorf("order already paid")
	}

	orderStatus, err := entities.SetStatus(status)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = s.paymentRepo.UpdatePayment(ctx, uuidID, orderStatus.Status, orderStatus.StatusDescription)

	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %w", err)
	}

	return orderStatus, nil
}

func (s *Payment) GenerateQRCode(ctx context.Context, id string) ([]byte, error) {

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	order, err := s.orderUseCases.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	if order.Status != 0 {
		return nil, fmt.Errorf("order already paid")
	}

	png, err := qrcode.Encode(id, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}
	return png, nil
}
