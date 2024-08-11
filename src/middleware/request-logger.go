package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var requestLog = logrus.New()

func RequestLogger(c *fiber.Ctx) error {
	requestLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	requestLog.Infof("Request: %s %s", c.Method(), c.Path())

	return c.Next()
}
