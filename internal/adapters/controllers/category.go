package controllers

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type CategoryController struct {
	usecase interfaces.CategoryUseCase
}

func NewCategoryController(u interfaces.CategoryUseCase) *CategoryController {
	return &CategoryController{
		usecase: u,
	}
}

func (c *CategoryController) CreateCategory(ctx context.Context, categoryDto dto.CreateCategoryRequest) (*entities.Category, error) {
	return c.usecase.CreateCategory(ctx, categoryDto)
}

func (c *CategoryController) ReplaceCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error) {
	return c.usecase.ReplaceCategory(ctx, id, category)
}

func (c *CategoryController) UpdateCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error) {
	return c.usecase.UpdateCategory(ctx, id, category)
}

func (c *CategoryController) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	return c.usecase.GetCategoryByID(ctx, id)
}

func (c *CategoryController) GetCategories(ctx context.Context, page int, size int) ([]entities.Category, error) {
	return c.usecase.GetCategories(ctx, page, size)
}

func (c *CategoryController) DeleteCategory(ctx context.Context, id string) error {
	return c.usecase.DeleteCategory(ctx, id)
}
