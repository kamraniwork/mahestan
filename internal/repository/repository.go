package repository

import (
	"context"
	"mahestan/internal/models"
)

type CommonBehaviourRepository[T models.Entity] interface {
	Create(ctx context.Context, ent T) (T, error)
	ReadAll(ctx context.Context) ([]T, error)
	UpdateByField(ctx context.Context, ent models.Entity, new_ent models.Entity) error
	Delete(ctx context.Context, ent models.Entity, id uint) error
	FindByID(ctx context.Context, id uint) (T, error)
	FindByField(ctx context.Context, fieldName string, value interface{}) ([]T, error)
	Count(ctx context.Context, ent models.Entity) (int64, error)
	Exists(ctx context.Context, fieldName string, value interface{}) (bool, error)
}

type CategoryRepository interface {
	CommonBehaviourRepository[models.Category]
	FindByName(ctx context.Context, name string) (*models.Category, error)
	FindByProductCount(ctx context.Context, count int) ([]models.Category, error)
	FindEmptyCategories(ctx context.Context) ([]models.Category, error)
	GetProducts(ctx context.Context, categoryID uint) ([]models.Product, error)
}
