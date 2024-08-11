package customvalidation

import "github.com/go-playground/validator/v10"

func RequiredName(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) >= 5
}
