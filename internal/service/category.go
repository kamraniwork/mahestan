package service

import (
	"context"
	"mahestan/internal/controller/serializers"
	"mahestan/internal/models"
	"mahestan/internal/repository"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(rep repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: rep,
	}
}

func (c *CategoryService) Create(ctx context.Context, input serializers.InputCategorySerializer) (models.Category, error) {
	ent, err := c.categoryRepository.Create(ctx, models.Category{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return ent, err
	}
	return ent, nil
}
