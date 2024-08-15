package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewSwagger() func(*fiber.Ctx) error {
	return swagger.New(swagger.Config{
		URL:          "/docs/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
		Title:        "Go Fiber Setup",
	})
}
