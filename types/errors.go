package types

import "fmt"

type (
	InvalidTypeError struct {
		receivedString     string
		expectedStringType string
	}

	ValidatorAddressUnmatchedError struct {
		hexAddress    string
		bech32Address string
	}

	ExceedMaxlengthError struct {
		variableName string
		length       string
		maxLength    string
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

func NewValidatorAddressUnmatched(hexAddress string, bech32Address string) ValidatorAddressUnmatchedError {
	return ValidatorAddressUnmatchedError{
		hexAddress:    hexAddress,
		bech32Address: bech32Address,
	}
}

func (e ValidatorAddressUnmatchedError) Error() string {
	return fmt.Sprintf("HEX " + e.hexAddress + " and Bech32 " + e.bech32Address + " do not match")
}

func NewExceedMaxlengthError(variableName string, length string, maxLength string) ExceedMaxlengthError {
	return ExceedMaxlengthError{
		variableName: variableName,
		length:       length,
		maxLength:    maxLength,
	}
}

func (e ExceedMaxlengthError) Error() string {
	return fmt.Sprintf("Variable " + e.variableName + " length " + e.length + " has exceed max length " + e.maxLength)
}
