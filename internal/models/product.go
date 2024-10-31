package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `gorm:"size:255;not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	SKU         string         `gorm:"size:100;unique;not null" json:"sku"`
	Stock       int            `gorm:"not null" json:"stock"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
	Reviews     []Review       `gorm:"foreignKey:ProductID"`
}

func (p Product) EntityName() string {
	return fmt.Sprint("Product_%s", p.Name)
}
