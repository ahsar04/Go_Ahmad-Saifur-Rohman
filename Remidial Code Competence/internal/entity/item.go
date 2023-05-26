package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	*gorm.Model

	ItemName     	string `json:"item_name" form:"item_name" validate:"required"`
	Description     string `json:"description" form:"description" validate:"required"`
	Stock        	string `json:"stock" form:"stock" validate:"required"`
	Price        	string `json:"price" form:"price" validate:"required"`
	CategoryId       string `json:"category_id" form:"category_id" validate:"required"`
	Category        Category `gorm:"foreignKey:CategoryId"`
}
