package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	// Logic to get a user
	return c.JSON(fiber.Map{"message": "Get User"})
}

func PostUser(c *fiber.Ctx) error {
	// Logic to create a user
	return c.JSON(fiber.Map{"message": "Post User"})
}
