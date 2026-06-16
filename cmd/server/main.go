package main

import (
	"log"

	"fmt"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)
	log.Println("servver started on port 3000")

	log.Fatal(app.Listen(":3000"))

	age, _ := service.CalculateAge("1990-05-10")

	fmt.Println(age)

}
