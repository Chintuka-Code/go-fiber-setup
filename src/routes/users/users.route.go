package users

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/controller"
	"struct-validation/src/validation"
)

func InitUserRoutes(group fiber.Router) {
	userGroup := group.Group("/user")
	userGroup.Get("/", controller.GetUser)
	userGroup.Post("/", validation.ValidatePostUser, controller.PostUser)
}
