package validation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PostUser struct {
	Name string `validate:"required,min=5,max=20"`
	Age  int    `validate:"required,teener"`
}

func ValidatePostUser(c *fiber.Ctx) error {
	var data PostUser
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := ValidateStruct(data)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).Send([]byte("Some Error"))
	}
	return c.Next()
}
