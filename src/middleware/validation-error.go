package middleware

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"struct-validation/src/shared"
)

func ValidationErrorCatch(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		if validationErrs, ok := err.(shared.ValidationErrorList); ok {
			errorDataMessage := validationErrs.Pretty()
			structName := color.New(color.FgRed).Sprint("ValidationErrorList")
			logrus.WithFields(logrus.Fields{
				"details": errorDataMessage,
			}).Error(fmt.Sprintf("[%s]", structName))

			return c.Status(400).JSON(fiber.Map{
				"error": errorDataMessage,
			})
		}

		return err
	}

	return nil
}
