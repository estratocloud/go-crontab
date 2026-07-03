package internal

import (
	"fmt"
)

type Field struct {
	value string
	err   error
}

// Error Check if there were any invalid values (eg hour=35)
func (f Field) Error() error {
	return f.err
}

// Format Get the field's string value (eg "*/5")
// or an error for invalid values (eg hour=35)
func (f Field) Format() (string, error) {
	return f.value, f.err
}

// MustFormat Similar to Format() but will panic if any invalid values were provided,
// ONLY use this method with known safe values, if you need to generate expressions
// from user input then use Format() and handle errors properly.
func (f Field) MustFormat() string {
	if f.err != nil {
		panic(fmt.Errorf("unhandled error, use Format() to handle: %w", f.err))
	}
	return f.value
}

// Every Create a Field instance that represents every possible value ("*")
func Every() Field {
	return Field{value: "*"}
}
