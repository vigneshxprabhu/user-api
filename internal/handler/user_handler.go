package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return c.JSON([]string{"BOB", "ALICE"})

}
