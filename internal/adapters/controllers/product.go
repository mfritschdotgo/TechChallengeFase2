package controllers

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type ProductController struct {
	usecase   interfaces.ProductUseCase
	presenter interfaces.ProductPresenter
}

func NewProductController(u interfaces.ProductUseCase, p interfaces.ProductPresenter) *ProductController {
	return &ProductController{
		usecase:   u,
		presenter: p,
	}
}

func (c *ProductController) CreateProduct(ctx context.Context, dto dto.CreateProductRequest) (*entities.Product, error) {
	return c.usecase.CreateProduct(ctx, dto)
}

func (c *ProductController) ReplaceProduct(ctx context.Context, id string, dto dto.CreateProductRequest) (*entities.Product, error) {
	return c.usecase.ReplaceProduct(ctx, id, dto)
}

func (c *ProductController) UpdateProduct(ctx context.Context, id string, dto dto.CreateProductRequest) (*entities.Product, error) {
	return c.usecase.UpdateProduct(ctx, id, dto)
}

func (c *ProductController) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	return c.usecase.GetProductByID(ctx, id)
}

func (c *ProductController) GetProducts(ctx context.Context, category string, page, size int) ([]entities.Product, error) {
	return c.usecase.GetProducts(ctx, category, page, size)
}

func (c *ProductController) DeleteProduct(ctx context.Context, id string) error {
	return c.usecase.DeleteProduct(ctx, id)
}
