package models

import "errors"

// ErrorInvalidURL is a global variable that holds an error.
// This error is returned when an invalid URL is encountered in the application.
var (
	ErrorInvalidURL = errors.New("invalid URL") // Error message for invalid URL
)
