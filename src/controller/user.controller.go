package controller

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/localize"
)

func GetUser(c *fiber.Ctx) error {
	message := localize.GetTranslation(c, "HELLO")
	return c.JSON(fiber.Map{"message": message})
}

func PostUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Post User"})
}
