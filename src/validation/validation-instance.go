package validation

import (
	"github.com/go-playground/validator/v10"

	"struct-validation/src/shared"
)

var Validator *validator.Validate

func InitValidator() *validator.Validate {
	Validator = validator.New()
	return Validator
}

func ValidateStruct(data interface{}) error {
	var validationErrors []shared.ValidationError
	errs := Validator.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, shared.ValidationError{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				HasError:    true,
			})
		}
		return shared.NewValidationError(validationErrors)
	}
	return nil
}
