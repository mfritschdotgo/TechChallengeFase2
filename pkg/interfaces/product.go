package interfaces

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
	GetProducts(ctx context.Context, categoryId string, page, limit int) ([]entities.Product, error)
	ReplaceProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
