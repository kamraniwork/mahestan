package repository

import (
	"context"
	"gorm.io/gorm"
	"mahestan/internal/models"
)

type CategoriesRepository struct {
	*CommonBehaviour[models.Category]
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{
		NewCommonBehaviour[models.Category](db),
	}
}

func (c *CategoriesRepository) FindByName(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category
	err := c.db.WithContext(ctx).Where("name = ?", name).First(&category).Error
	return &category, err
}

func (c *CategoriesRepository) FindByProductCount(ctx context.Context, count int) ([]models.Category, error) {
	var categories []models.Category
	err := c.db.WithContext(ctx).Joins("JOIN products ON products.category_id = categories.id").
		Group("categories.id").Having("COUNT(products.id) = ?", count).Find(&categories).Error
	return categories, err
}

func (c *CategoriesRepository) FindEmptyCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := c.db.WithContext(ctx).Joins("LEFT JOIN products ON products.category_id = categories.id").
		Where("products.id IS NULL").Find(&categories).Error
	return categories, err
}

func (c *CategoriesRepository) GetProducts(ctx context.Context, categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := c.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}
