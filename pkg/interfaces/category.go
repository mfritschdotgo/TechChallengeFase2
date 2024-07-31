package interfaces

import (
	"context"

	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/pkg/dto"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*entities.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error)
	ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
}

type CategoryGateway interface {
	CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*entities.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]entities.Category, error)
	ReplaceCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	UpdateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, dto dto.CreateCategoryRequest) (*entities.Category, error)
	ReplaceCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error)
	UpdateCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*entities.Category, error)
	GetCategories(ctx context.Context, page, size int) ([]entities.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	InitializeCategories(ctx context.Context) error
}
