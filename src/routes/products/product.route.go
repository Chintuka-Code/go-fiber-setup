package products

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/controller"
)

func InitProductRoutes(group fiber.Router) {
	productGroup := group.Group("/product")
	productGroup.Get("/", controller.GetProduct)
	productGroup.Post("/", controller.PostProduct)
}
