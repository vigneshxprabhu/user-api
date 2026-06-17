package handler

import (
	//"fmt"
	"strconv"
	"user-api/internal/models"
	"user-api/internal/repository"
	//"user-api/internal/service"
	"time"

	"github.com/gofiber/fiber/v2"
	database "user-api/db/sqlc/generated"
)

var users = []models.User{
	{
		ID:   1,
		Name: "Alice",
		DOB:  "1990-05-10",
	},
	{
		ID:   2,
		Name: "Bob",
		DOB:  "1995-08-20",
	},
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {

	users, err := h.repo.ListUsers()

	if err != nil {
		return c.Status(500).SendString("Failed to fetch users")
	}

	return c.JSON(users)

}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid ID")
	}

	user, err := h.repo.GetUser(int32(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).
			SendString("User not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid request body")
	}

	dob, err := time.Parse("2006-01-02", req.DOB)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid DOB format")
	}

	user, err := h.repo.CreateUser(
		database.CreateUserParams{
			Name: req.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Failed to create user")
	}

	return c.Status(fiber.StatusCreated).
		JSON(user)
}
func UpdateUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}
	var req models.CreateUserRequest

	err = c.BodyParser(&req)

	if err != nil {
		return c.Status(400).SendString("Invalid Request")
	}
	if req.Name == "" || req.DOB == "" {
		return c.Status(400).SendString("Name and DOB are required")
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = req.Name
			users[i].DOB = req.DOB

			return c.JSON(users[i])
		}

	}
	return c.Status(404).SendString("User not found")
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	for i, user := range users {

		if user.ID == id {

			users = append(users[:i], users[i+1:]...)

			return c.SendStatus(204)
		}
	}
	return c.Status(404).SendString("User not found")

}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}
