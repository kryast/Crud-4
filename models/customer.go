package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"omitempty,max=15"`
}
