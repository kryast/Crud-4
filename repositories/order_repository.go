package repositories

import (
	"github.com/kryast/Crud-4.git/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	FindAll() ([]models.Order, error)
	FindByID(id uint) (models.Order, error)
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

func (r *orderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Customer").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) FindByID(id uint) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Customer").First(&order, id).Error
	return order, err
}
