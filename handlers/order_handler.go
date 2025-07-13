package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-4.git/models"
	"github.com/kryast/Crud-4.git/services"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{s}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, order)
}
