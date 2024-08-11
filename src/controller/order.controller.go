package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get Order"})
}

func PostOrder(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Post Order"})
}
