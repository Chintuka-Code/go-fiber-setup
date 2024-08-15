package validation

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type File struct {
	Filename   string
	Size       int64
	FileHeader *multipart.FileHeader
}

type PostUser struct {
	Name    string   `validate:"required,min=5,max=20"`
	Age     int      `validate:"required"`
	Address []string `validate:"required,dive,required"`
	Hobbies []string `validate:"required,dive,uuid"`
	Image   File     `validate:"required,file,image"`
}

func ValidatePostUser(c *fiber.Ctx) error {
	var data PostUser
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := ValidateStruct(data)
	if err != nil {
		return err
	}
	return c.Next()
}
