package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-4.git/handlers"
)

func SetupRouter(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)

	return r
}
