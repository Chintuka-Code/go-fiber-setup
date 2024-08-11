package customvalidation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidation(validator *validator.Validate) {
	err := validator.RegisterValidation("requiredName", RequiredName)

	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}
}
