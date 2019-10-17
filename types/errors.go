package types

import "fmt"

type (
	InvalidTypeError struct {
		receivedString     string
		expectedStringType string
	}

	ValidatorAddressUnmatched struct {
		hexAddress    string
		bech32Address string
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

func NewValidatorAddressUnmatched(hexAddress string, bech32Address string) ValidatorAddressUnmatched {
	return ValidatorAddressUnmatched{
		hexAddress:    hexAddress,
		bech32Address: bech32Address,
	}
}

func (e ValidatorAddressUnmatched) Error() string {
	return fmt.Sprintf("HEX " + e.hexAddress + " and Bech32 " + e.bech32Address + " do not match")
}
