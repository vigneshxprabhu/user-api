package routes

import (
	"user-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", handler.GetUsers)
	app.Get("/users/:id", handler.GetUserByID)
	app.Post("/users", handler.CreateUser)
	app.Put("/users/:id", handler.UpdateUser)
}
