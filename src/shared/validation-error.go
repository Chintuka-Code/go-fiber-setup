package shared

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
