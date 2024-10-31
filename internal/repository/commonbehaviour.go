package repository

import (
	"context"
	"gorm.io/gorm"
	"mahestan/internal/models"
)

type CommonBehaviour[T models.Entity] struct {
	db *gorm.DB
}

func NewCommonBehaviour[T models.Entity](db *gorm.DB) *CommonBehaviour[T] {
	return &CommonBehaviour[T]{
		db: db,
	}
}

func (c *CommonBehaviour[T]) Create(ctx context.Context, ent T) (T, error) {
	err := c.db.WithContext(ctx).Create(&ent).Error
	return ent, err
}

func (c *CommonBehaviour[T]) ReadAll(ctx context.Context) ([]T, error) {
	var commonObject []T
	err := c.db.WithContext(ctx).Find(&commonObject).Error
	return commonObject, err
}

func (c *CommonBehaviour[T]) UpdateByField(ctx context.Context, ent models.Entity, newEnt models.Entity) error {
	return c.db.WithContext(ctx).
		Model(&ent).
		Updates(newEnt).
		Error
}

func (c *CommonBehaviour[T]) Delete(ctx context.Context, ent models.Entity, id uint) error {

	return c.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&ent).
		Error
}

func (c *CommonBehaviour[T]) FindByID(ctx context.Context, id uint) (T, error) {
	var common T
	err := c.db.WithContext(ctx).First(&common, id).Error
	return common, err
}

func (c *CommonBehaviour[T]) FindByField(ctx context.Context, fieldName string, value interface{}) ([]T, error) {
	var common []T
	err := c.db.WithContext(ctx).Where(fieldName+" = ?", value).Find(&common).Error
	return common, err
}

func (c *CommonBehaviour[T]) Count(ctx context.Context, ent models.Entity) (int64, error) {
	var count int64
	err := c.db.WithContext(ctx).Model(&ent).Count(&count).Error
	return count, err
}

func (c *CommonBehaviour[T]) Exists(ctx context.Context, fieldName string, value interface{}) (bool, error) {
	var exists bool
	err := c.db.WithContext(ctx).Model(&models.Category{}).Select("count(*) > 0").Where(fieldName+" = ?", value).Find(&exists).Error
	return exists, err
}
