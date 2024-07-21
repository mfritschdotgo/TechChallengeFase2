package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallengefase2/internal/adapters/dto"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/entities"
	"github.com/mfritschdotgo/techchallengefase2/internal/domain/interfaces"
)

type Category struct {
	categoryRepo interfaces.CategoryRepository
}

func NewCategory(repo interfaces.CategoryRepository) *Category {
	return &Category{
		categoryRepo: repo,
	}
}

func (s *Category) CreateCategory(ctx context.Context, dto dto.CreateCategoryRequest) (*entities.Category, error) {
	category, err := entities.NewCategory(dto.Name, dto.Description)
	if err != nil {
		return nil, err
	}

	_, err = s.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Category) ReplaceCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	categoryDto, err := s.GetCategoryByID(ctx, uuidID.String())

	if err != nil {
		return nil, err
	}

	categoryDto.Name = category.Name
	categoryDto.Description = category.Description
	categoryDto.UpdatedAt = time.Now()

	if _, err = s.categoryRepo.ReplaceCategory(ctx, categoryDto); err != nil {
		return nil, fmt.Errorf("failed to replace category: %w", err)
	}

	return categoryDto, nil
}

func (s *Category) UpdateCategory(ctx context.Context, id string, category *entities.Category) (*entities.Category, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	category.ID = uuidID
	category.UpdatedAt = time.Now()

	if _, err = s.categoryRepo.UpdateCategory(ctx, category); err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}
	response, err := s.GetCategoryByID(ctx, uuidID.String())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Category) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	category, err := s.categoryRepo.GetCategoryByID(ctx, uuidID.String())
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	return category, nil
}

func (s *Category) GetCategories(ctx context.Context, page, size int) ([]entities.Category, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	categories, err := s.categoryRepo.GetCategories(ctx, page, size)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Category) DeleteCategory(ctx context.Context, id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	if err = s.categoryRepo.DeleteCategory(ctx, uuidID.String()); err != nil {
		return fmt.Errorf("category not found or error deleting category: %w", err)
	}

	return nil
}

func (s *Category) InitializeCategories(ctx context.Context) error {
	categories, err := s.GetCategories(ctx, 1, 1)
	if err != nil {
		return err
	}

	if len(categories) == 0 {
		initialCategories := []dto.CreateCategoryRequest{
			{Name: "Lanche", Description: "Categoria de Lanches"},
			{Name: "Acompanhamento", Description: "Categoria de Acompanhamentos"},
			{Name: "Bebida", Description: "Categoria de Bebidas"},
			{Name: "Sobremesa", Description: "Categoria de Sobremesas"},
		}

		for _, cat := range initialCategories {
			_, err := s.CreateCategory(ctx, cat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
