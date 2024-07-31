package gateways

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/interfaces"
)

type CategoryGateway struct {
	repo interfaces.CategoryRepository
}

func NewCategoryGateway(repo interfaces.CategoryRepository) *CategoryGateway {
	return &CategoryGateway{
		repo: repo,
	}
}

func (g *CategoryGateway) CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	return g.repo.CreateCategory(ctx, category)
}

func (g *CategoryGateway) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	return g.repo.GetCategoryByID(ctx, id)
}

func (g *CategoryGateway) ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	return g.repo.ReplaceCategory(ctx, category)
}

func (g *CategoryGateway) UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	return g.repo.UpdateCategory(ctx, category)
}

func (g *CategoryGateway) DeleteCategory(ctx context.Context, id string) error {
	return g.repo.DeleteCategory(ctx, id)
}

func (g *CategoryGateway) GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error) {
	return g.repo.GetCategories(ctx, page, limit)
}
