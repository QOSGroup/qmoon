package types

import "fmt"

type (
	InvalidTypeError struct {
		receivedString     string
		expectedStringType string
	}
)

// New returns an error that formats as the given text.
func NewError(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func NewInvalidTypeError(received string, expectedType string) InvalidTypeError {
	return InvalidTypeError{
		receivedString:     received,
		expectedStringType: expectedType,
	}
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("Wrong type: " + e.receivedString + " is not a " + e.expectedStringType)
}
