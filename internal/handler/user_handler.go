package handler

import (
	//"fmt"
	"strconv"
	"user-api/internal/models"
	"user-api/internal/repository"

	"time"
	"user-api/internal/service"

	database "user-api/db/sqlc/generated"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {

	users, err := h.repo.ListUsers()

	if err != nil {
		return c.Status(500).SendString("Failed to fetch users")
	}

	var response []models.UserResponse

	for _, user := range users {
		age, err := service.CalculateAge(
			user.Dob.Format("2006-01-02"),
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("Failed to calculate age")
		}
		response = append(response, models.UserResponse{
			ID:   int(user.ID),
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  age,
		})
	}
	return c.JSON(response)
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

	age, err := service.CalculateAge(
		user.Dob.Format("2006-01-02"),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Failed to calculate age")
	}

	response := models.UserResponse{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}

	return c.JSON(response)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid request body")
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Validation failed")
	}
	dob, err := time.Parse("2006-01-02", req.DOB)

	if dob.After(time.Now()) {
		return c.Status(fiber.StatusBadRequest).
			SendString("DOB  cannot be in the future")
	}

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

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid ID")
	}

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
	user, err := h.repo.UpdateUser(
		database.UpdateUserParams{
			ID:   int32(id),
			Name: req.Name,
			Dob:  dob,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Failed to update user")
	}
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid ID")
	}
	err = h.repo.DeleteUser(int32(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Delete failed")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}
