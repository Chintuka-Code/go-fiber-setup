package customvalidation

import (
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidation(validator *validator.Validate) {
	validator.RegisterValidation("requiredName", RequiredName)
}
