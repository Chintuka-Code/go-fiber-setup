package main

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", hello)

	app.Get("/api/user", user.GetUser)
}

func main() {
	// Fiber instance
	app := fiber.New()
	setUpRoutes(app)
	// start server
	app.Listen(":8000")
}
