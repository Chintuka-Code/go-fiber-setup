package middleware

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/shared"
)

func ValidationErrorCatch(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		if validationErrs, ok := err.(shared.ValidationErrorList); ok {
			errorDataMessage := validationErrs.Pretty()
			return c.Status(400).JSON(errorDataMessage)
		}
		return err
	}

	return nil
}
