// cmd/main.go
package main

import (
	"go-auth/database"
	"go-auth/internal/handler"
	"go-auth/internal/model"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	database.DB.AutoMigrate(&model.User{})

	app := fiber.New()

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	app.Listen(":3000")
}
