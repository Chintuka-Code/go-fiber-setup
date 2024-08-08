package user

import (
	"struct-validation/src/exception"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	return c.SendString("All Bookmarks")
}

func PostUser(c *fiber.Ctx) error {
	return exception.New(fiber.StatusBadRequest, "Failed to create user")
}
