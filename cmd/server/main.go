package main

import (
	"log"

	"user-api/config"
	database "user-api/db/sqlc/generated"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/repository"
	"user-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()

	defer db.Close()
	queries := database.New(db)
	repo := repository.NewUserRepository(queries)

	userHandler := handler.NewUserHandler(repo)

	_ = queries

	logger.Init()

	app := fiber.New()
	app.Use(middleware.RequestID)
	app.Use(middleware.RequestLogger)
	routes.SetupRoutes(app, userHandler)
	log.Println("server started on port 8080")

	log.Fatal(app.Listen(":8080"))

}
