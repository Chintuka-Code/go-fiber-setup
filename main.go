package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"struct-validation/src/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", hello)

	app.Get("/api/user", user.GetUser)
	app.Post("/api/user", user.PostUser)

}

func main() {
	// Fiber instance
	app := fiber.New()
	setUpRoutes(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
