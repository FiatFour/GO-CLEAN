package main

import (
	"log"

	// "github.com/fiatfour/go-clean/adapters"
	"github.com/fiatfour/go-clean/adapters"
	"github.com/fiatfour/go-clean/entities"
	"github.com/fiatfour/go-clean/usecases"

	// "github.com/fiatfour/go-clean/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database %v", err)
	}

	db.Migrator().CreateTable(&entities.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen("localhost:8080")
}
