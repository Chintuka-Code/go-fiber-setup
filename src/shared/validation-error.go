package shared

import "fmt"

type ValidationError struct {
	FailedField string      `json:"field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
	HasError    bool        `json:"error"` // Renamed field to avoid conflict
}

// Implementing the error interface
func (v *ValidationError) Error() string {
	return v.FailedField + " failed on the " + v.Tag + " validation"
}

type ValidationErrorList []ValidationError

func (v ValidationErrorList) Error() string {
	return "Validation failed"
}

func (v ValidationErrorList) Errors() []ValidationError {
	return v
}

func NewValidationError(errors []ValidationError) ValidationErrorList {
	return ValidationErrorList(errors)
}

func (v ValidationErrorList) Pretty() map[string][]string {
	prettyErrors := make(map[string][]string)
	for _, err := range v {
		if err.HasError {
			prettyErrors[err.FailedField] = append(prettyErrors[err.FailedField], fmt.Sprintf("Field '%s' failed validation with tag '%s' for value '%v'", err.FailedField, err.Tag, err.Value))
		}
	}
	return prettyErrors
}
