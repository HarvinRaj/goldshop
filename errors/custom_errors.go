package errors

import "fmt"

type MyError struct {
	code    int
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func New(code int) *MyError {
	message, exists := ErrorMessages[code]
	if !exists {
		message = "Error code not found in custom errors"
	}

	return &MyError{
		code:    code,
		message: message,
	}
}
