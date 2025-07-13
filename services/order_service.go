package services

import "github.com/kryast/Crud-4.git/repositories"

type OrderService interface {
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(r repositories.OrderRepository) OrderService {
	return &orderService{r}
}
