package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

func (c Category) EntityName() string {
	return fmt.Sprint("Category_%s", c.Name)
}
