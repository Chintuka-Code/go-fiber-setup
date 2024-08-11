package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Product"})
}

func PostProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Post Product"})
}
