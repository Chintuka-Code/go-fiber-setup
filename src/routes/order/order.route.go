package order

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/controller"
)

func InitOrderRoutes(group fiber.Router) {
	orderGroup := group.Group("/orders")
	orderGroup.Get("/", controller.GetOrder)
	orderGroup.Post("/", controller.PostOrder)
}
