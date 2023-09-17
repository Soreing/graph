package resource

// ValidationError is an error type returned when a resource failed to validate
// its integrity due to it reading the wrong data type or one of its properties
// failing to validate
type ValidationError struct{}

// NewValidationError creates a new ValidationError
func NewValidationError() error {
	return &ValidationError{}
}

// Error returns an error message as string
func (ValidationError) Error() string {
	return "validation failed"
}
