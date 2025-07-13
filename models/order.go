package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint     `json:"customer_id" binding:"required"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`
	Item       string   `json:"item" binding:"required"`
	Quantity   int      `json:"quantity" binding:"required,min=1"`
	Status     string   `json:"status" binding:"required,oneof=pending paid"`
}
