package services

import (
	"github.com/kryast/Crud-4.git/models"
	"github.com/kryast/Crud-4.git/repositories"
)

type OrderService interface {
	CreateOrder(order *models.Order) error
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
