package middleware

import (
	"time"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	methodColor = color.New(color.FgCyan).SprintFunc()
	urlColor    = color.New(color.FgGreen).SprintFunc()
	ipColor     = color.New(color.FgYellow).SprintFunc()
	statusColor = color.New(color.FgMagenta).SprintFunc()
	timeColor   = color.New(color.FgBlue).SprintFunc()
)

func RequestLogger(c *fiber.Ctx) error {
	LOG.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	startTime := time.Now()
	responseTime := time.Since(startTime).Milliseconds()
	LOG.Infof(
		"{ %s } %s | IP: %s | Status: %s | Time: %sms",
		methodColor(c.Method()),
		urlColor(c.Path()),
		ipColor(c.IP()),
		statusColor(c.Response().StatusCode()),
		timeColor(responseTime),
	)

	return c.Next()
}
