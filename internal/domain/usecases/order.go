package usecases

import (
	"context"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type Order struct {
	orderRepo       interfaces.OrderRepository
	clientUseCases  *Client
	productUseCases *Product
}

func NewOrder(repo interfaces.OrderRepository, client *Client, product *Product) *Order {
	return &Order{
		orderRepo:       repo,
		clientUseCases:  client,
		productUseCases: product,
	}
}

func (s *Order) CreateOrder(ctx context.Context, dto dto.CreateOrderRequest) (*entities.Order, error) {
	_, err := s.clientUseCases.GetClientByCPF(ctx, dto.Client)
	if err != nil {
		return nil, fmt.Errorf("client validation failed: %w", err)
	}

	total := 0.0
	productDetails := make(map[string]struct {
		Price float64
		Name  string
	})

	for _, item := range dto.Products {
		product, err := s.productUseCases.GetProductByID(ctx, item.ID)
		if err != nil {
			return nil, fmt.Errorf("product validation failed for product ID %s: %w", item.ID, err)
		}
		productDetails[item.ID] = struct {
			Price float64
			Name  string
		}{
			Price: product.Price,
			Name:  product.Name,
		}
		total += product.Price * float64(item.Quantity)
	}

	total = math.Round(total*100) / 100
	items := ConvertDTOtoSlice(dto.Products, productDetails)

	order, err := entities.NewOrder(dto.Client, items, 0, total, "created")
	if err != nil {
		return nil, fmt.Errorf("failed to create order instance: %w", err)
	}

	savedOrder, err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return savedOrder, nil
}

func ConvertDTOtoSlice(dtoProducts []dto.ProductItem, productDetails map[string]struct {
	Price float64
	Name  string
}) []entities.OrderItem {
	var domainItems []entities.OrderItem
	for _, item := range dtoProducts {
		details := productDetails[item.ID]
		domainItems = append(domainItems, entities.OrderItem{
			ProductID:   item.ID,
			ProductName: details.Name,
			Quantity:    item.Quantity,
			Price:       math.Round((details.Price*float64(item.Quantity))*100) / 100,
		})
	}
	return domainItems
}

func (s *Order) GetOrderByID(ctx context.Context, id string) (*entities.Order, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	order, err := s.orderRepo.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	return order, nil
}

func (s *Order) GetOrders(ctx context.Context, page, size int) ([]entities.Order, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	orders, err := s.orderRepo.GetOrders(ctx, page, size)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Order) SetOrderStatus(ctx context.Context, id string, status int) (*entities.OrderStatus, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	_, err = s.orderRepo.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	orderStatus, err := entities.SetStatus(status)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = s.orderRepo.SetStatus(ctx, uuidID, orderStatus.Status, orderStatus.StatusDescription)

	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %w", err)
	}

	return orderStatus, nil
}
