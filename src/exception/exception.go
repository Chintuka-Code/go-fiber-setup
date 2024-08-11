package exception

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

type Exception struct {
	StatusCode int
	Message    string
	Err        error
}

func New(statusCode int, message string) *Exception {
	err := errors.New(message)
	return &Exception{
		StatusCode: statusCode,
		Message:    message,
		Err:        errors.WithStack(err),
	}
}

func (e *Exception) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s, StackTrace: %+v", e.StatusCode, e.Message, e.Err)
}

func (e *Exception) Unwrap() error {
	return e.Err
}

// GetStackTrace returns a filtered stack trace
func (e *Exception) GetStackTrace() string {
	stackTrace := fmt.Sprintf("%+v", e.Err)
	lines := strings.Split(stackTrace, "\n")

	// Filter out non-application lines and ensure uniqueness
	seen := make(map[string]bool)
	var filteredLines []string
	for _, line := range lines {
		if strings.Contains(line, "struct-validation/src") && !strings.Contains(line, "middleware.GlobalErrorCatch") {
			if !seen[line] {
				seen[line] = true
				filteredLines = append(filteredLines, line)
			}
		}
	}

	stack := strings.Join(filteredLines, "\n")
	// Using color for better CLI visibility
	return color.New(color.FgRed).Sprint(stack)
}
