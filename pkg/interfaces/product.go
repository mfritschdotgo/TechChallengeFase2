package interfaces

import (
	"context"
	"net/http"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
	ReplaceProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, categoryId string, page, limit int) ([]entities.Product, error)
}

type ProductGateway interface {
	CreateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
	ReplaceProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	UpdateProduct(ctx context.Context, product *entities.Product) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, categoryId string, page, limit int) ([]entities.Product, error)
}

type ProductUseCase interface {
	CreateProduct(ctx context.Context, dto dto.CreateProductRequest) (*entities.Product, error)
	ReplaceProduct(ctx context.Context, id string, dto dto.CreateProductRequest) (*entities.Product, error)
	UpdateProduct(ctx context.Context, id string, dto dto.CreateProductRequest) (*entities.Product, error)
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
	GetProducts(ctx context.Context, category string, page, size int) ([]entities.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type ProductPresenter interface {
	JSON(w http.ResponseWriter, data interface{}, statusCode int)
	Error(w http.ResponseWriter, message string, statusCode int)
}
