package main

import (
	"log"

	"context"
	"github.com/gofiber/fiber/v2"
	"user-api/config"
	database "user-api/db/sqlc/generated"
	"user-api/internal/handler"
	"user-api/internal/repository"
	"user-api/internal/routes"
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
	app := fiber.New()

	routes.SetupRoutes(app, userHandler)
	log.Println("servver started on port 3000")

	log.Fatal(app.Listen(":3000"))

}
