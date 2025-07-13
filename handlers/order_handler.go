package handlers

import "github.com/kryast/Crud-4.git/services"

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{s}
}
