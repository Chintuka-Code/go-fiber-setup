package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"struct-validation/src/exception"
)

var LOG = logrus.New()

// GlobalErrorCatch middleware
func GlobalErrorCatch(c *fiber.Ctx) error {
	err := c.Next()

	if err == nil {
		return nil
	}

	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	stackTrace := ""

	if ex, ok := err.(*exception.Exception); ok {
		code = ex.StatusCode
		message = ex.Message
		stackTrace = ex.GetStackTrace()
	}

	LOG.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	LOG.WithFields(logrus.Fields{
		"status_code": code,
		"error":       message,
	}).Error("Error occurred")
	LOG.Error("StackTrace:\n" + stackTrace)

	return c.Status(code).SendString(message)
}
