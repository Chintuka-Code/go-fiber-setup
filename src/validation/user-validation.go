package validation

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

// File represents a file uploaded by the user
// swagger:model File
type File struct {
	// Filename of the uploaded file
	// example: "profile.jpg"
	Filename string `json:"filename"`
	// Size of the uploaded file in bytes
	// example: 123456
	Size int64 `json:"size"`
	// FileHeader contains metadata about the file
	FileHeader *multipart.FileHeader `json:"-"`
}

// PostUser represents the data required to create a new user
// swagger:model PostUser
type PostUser struct {
	// Name of the user
	// example: "John Doe"
	Name string `json:"name" validate:"required,min=5,max=20"`
	// Age of the user
	// example: 30
	Age int `json:"age" validate:"required"`
	// Address of the user
	// example: ["123 Main St", "Apt 4B"]
	Address []string `json:"address" validate:"required,dive,required"`
	// Hobbies of the user
	// example: ["reading", "gaming"]
	Hobbies []string `json:"hobbies" validate:"required,dive,uuid"`
	// Image uploaded by the user
	Image multipart.File `json:"image" validate:"required,file,image"`
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
