package controller

import (
	"github.com/gofiber/fiber/v2"

	"struct-validation/src/localize"
)

// @Summary Get a user
// @Description Get user details
// @Tags user
// @Accept   multipart/form-data
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /user [get]
func GetUser(c *fiber.Ctx) error {
	message := localize.GetTranslation(c, "HELLO")
	return c.JSON(fiber.Map{"message": message})
}

// PostUserHandler handles the POST /user request
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags user
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "Name of the user"
// @Param age formData int true "Age of the user"
// @Param address formData []string true "Address of the user"
// @Param hobbies formData []string true "Hobbies of the user"
// @Param image formData file true "Image uploaded by the user"
// @Success 200 {object} map[string]string
// @Router /user [post]
func PostUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Post User"})
}
