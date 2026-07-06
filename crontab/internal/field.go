package internal

import (
	"fmt"
	"sort"
	"strconv"
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

func validate(name string, min int, max int, values ...int) error {
	if len(values) == 0 {
		return fmt.Errorf("no %s values were passed", name)
	}
	for _, value := range values {
		if value < min || value > max {
			return fmt.Errorf("invalid %s (%d) must be between %d and %d", name, value, min, max)
		}
	}
	return nil
}

// Every Create a Field instance that represents every possible value ("*")
func Every() Field {
	return Field{value: "*"}
}

// Every Create a Field instance for a single value ("0")
func Value(name string, min int, max int, value int) Field {
	err := validate(name, min, max, value)
	if err != nil {
		return Field{err: err}
	}
	return Field{value: strconv.Itoa(value)}
}

// List Create a Field instance for a list of values ("4,8,12")
func List(name string, min int, max int, values ...int) Field {
	err := validate(name, min, max, values...)
	if err != nil {
		return Field{err: err}
	}

	values = append([]int(nil), values...)
	sort.Ints(values)

	result := ""
	for _, value := range values {
		if result != "" {
			result += ","
		}
		result += strconv.Itoa(value)
	}

	return Field{value: result}
}

// Range Create a Field instance for a range of values ("3-5")
func Range(name string, min int, max int, from int, to int) Field {
	err := validate(name, min, max, from, to)
	if err != nil {
		return Field{err: err}
	}
	if from > to {
		return Field{err: fmt.Errorf("invalid %s range, 'from' (%d) must be less than 'to' (%d)", name, from, to)}
	}
	return Field{value: fmt.Sprintf("%d-%d", from, to)}
}

// Range Create a Field instance for a step of values ("*/5")
func EveryX(name string, max int, value int) Field {
	if value == 1 {
		return Field{value: "*"}
	}
	err := validate(name, 1, max, value)
	if err != nil {
		return Field{err: err}
	}
	return Field{value: fmt.Sprintf("*/%d", value)}
}
