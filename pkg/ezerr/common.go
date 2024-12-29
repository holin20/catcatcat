package ezerr

import (
	"fmt"
)

// WrapError wraps an existing error with a new message and returns it
func WrapError(err error, message string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", message, err)
	}
	return fmt.Errorf(message)
}
