package interfaces

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*entities.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error)
	ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
}
