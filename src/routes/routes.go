package routes

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/exception"
	"struct-validation/src/routes/order"
	"struct-validation/src/routes/products"
	"struct-validation/src/routes/users"
)

func InitRoutes(app *fiber.App, baseRoute string) {
	api := app.Group(baseRoute)

	// Initialize individual route groups
	users.InitUserRoutes(api)
	products.InitProductRoutes(api)
	order.InitOrderRoutes(api)

	app.All("*", func(c *fiber.Ctx) error {
		return exception.New(fiber.StatusNotFound, "Route not found")
	})
}
