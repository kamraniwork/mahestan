package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Rating    int     `gorm:"not null" json:"rating"`
	Comment   string  `gorm:"type:text" json:"comment"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (r Review) EntityName() string {
	return fmt.Sprint("Review_%v", r.ID)
}
