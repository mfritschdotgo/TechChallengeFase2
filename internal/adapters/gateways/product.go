package gateways

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type ProductGateway struct {
	repo interfaces.ProductRepository
}

func NewProductGateway(repo interfaces.ProductRepository) interfaces.ProductGateway {
	return &ProductGateway{
		repo: repo,
	}
}

func (g *ProductGateway) CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	return g.repo.CreateProduct(ctx, product)
}

func (g *ProductGateway) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	return g.repo.GetProductByID(ctx, id)
}

func (g *ProductGateway) ReplaceProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	return g.repo.ReplaceProduct(ctx, product)
}

func (g *ProductGateway) UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	return g.repo.UpdateProduct(ctx, product)
}

func (g *ProductGateway) DeleteProduct(ctx context.Context, id string) error {
	return g.repo.DeleteProduct(ctx, id)
}

func (g *ProductGateway) GetProducts(ctx context.Context, categoryId string, page, limit int) ([]entities.Product, error) {
	return g.repo.GetProducts(ctx, categoryId, page, limit)
}
