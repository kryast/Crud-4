package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-4.git/handlers"
)

func SetupRouter(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.GetOrders)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)

	return r
}
