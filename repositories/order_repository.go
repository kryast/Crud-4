package repositories

import (
	"github.com/kryast/Crud-4.git/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}
