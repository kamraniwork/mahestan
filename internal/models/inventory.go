package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (i Inventory) EntityName() string {
	return fmt.Sprint("Inventory_%v", i.ID)
}
