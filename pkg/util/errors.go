package util

import "fmt"

// Represents a validation error.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Message)
}

// Represents a not found error.
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %s", e.Message)
}

// Represents a conflict error.
type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("conflict error: %s", e.Message)
}

// Represents an internal error.
type InternalError struct {
	Message string
	Err     error // The error that caused this error.
}

func (e *InternalError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("internal server error: %s: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("internal server error: %s", e.Message)
}
