package models

import (
	"fmt"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	URL       string `gorm:"size:255;not null" json:"url"`
	ProductID uint   `gorm:"not null" json:"product_id"`
}

func (pi ProductImage) EntityName() string {
	return fmt.Sprint("Product_images_%v", pi.ID)
}
