package users

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/controller"
)

func InitUserRoutes(group fiber.Router) {
	userGroup := group.Group("/user")
	userGroup.Get("/", controller.GetUser)
	userGroup.Post("/", controller.PostUser)
}
