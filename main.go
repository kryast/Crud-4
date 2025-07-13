package main

import (
	"log"

	"github.com/kryast/Crud-4.git/database"
	"github.com/kryast/Crud-4.git/handlers"
	"github.com/kryast/Crud-4.git/models"
	"github.com/kryast/Crud-4.git/repositories"
	"github.com/kryast/Crud-4.git/router"
	"github.com/kryast/Crud-4.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}

	db.AutoMigrate(&models.Customer{}, &models.Order{})

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := router.SetupRouter(orderHandler)
	r.Run(":8080")
}
