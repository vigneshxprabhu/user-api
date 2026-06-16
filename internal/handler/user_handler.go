package handler

import (
	//"fmt"
	"user-api/internal/models"

	"github.com/gofiber/fiber/v2"

	"strconv"
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

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)

}

func GetUserByID(c *fiber.Ctx) error {
	// users := []models.User{
	// 	{
	// 		ID:   1,
	// 		Name: "Alice",
	// 		DOB:  "1990-05-10",
	// 	},
	// 	{
	// 		ID:   2,
	// 		Name: "Bob",
	// 		DOB:  "1995-08-20",
	// 	},
	// }
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	for _, user := range users {
		if id == user.ID {
			return c.JSON(user)
		}

	}
	return c.Status(404).SendString("User not found")

}

func CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(400).SendString("Invalid Request")
	}

	if req.Name == "" || req.DOB == "" {
		return c.Status(400).SendString("Name and DOB are required")
	}

	newUser := models.User{
		ID:   len(users) + 1,
		Name: req.Name,
		DOB:  req.DOB,
	}

	users = append(users, newUser)

	return c.Status(201).JSON(newUser)
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
