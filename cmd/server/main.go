package main

import (
	"log"

	"context"
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
	users, err := queries.ListUsers(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	log.Println(users)
	_ = queries

	logger.Init()

	app := fiber.New()
	app.Use(middleware.RequestID)
	app.Use(middleware.RequestLogger)
	routes.SetupRoutes(app, userHandler)
	log.Println("servver started on port 3000")

	log.Fatal(app.Listen(":3000"))

}
