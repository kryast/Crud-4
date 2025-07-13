package services

import (
	"github.com/kryast/Crud-4.git/models"
	"github.com/kryast/Crud-4.git/repositories"
)

type OrderService interface {
	CreateOrder(order *models.Order) error
	GetOrders() ([]models.Order, error)
	GetOrderByID(id uint) (models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(r repositories.OrderRepository) OrderService {
	return &orderService{r}
}

func (s *orderService) CreateOrder(order *models.Order) error {
	return s.repo.Create(order)
}

func (s *orderService) GetOrders() ([]models.Order, error) {
	return s.repo.FindAll()
}

func (s *orderService) GetOrderByID(id uint) (models.Order, error) {
	return s.repo.FindByID(id)
}

func (s *orderService) UpdateOrder(order *models.Order) error {
	return s.repo.Update(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
