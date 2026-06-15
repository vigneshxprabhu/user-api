package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"user-api/internal/routes"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)
	log.Println("servver started on port 3000")

	log.Fatal(app.Listen(":3000"))
}
