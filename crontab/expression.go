package crontab

import (
	"fmt"

	"github.com/estratocloud/go-crontab/crontab/internal"
)

type Expression struct {
	Minute  internal.Field
Hour    internal.Field
	Day     internal.Field
	Month   internal.Field
	Weekday internal.Field
}

func NewExpression(e *Expression) *Expression {
	if e == nil {
		e = &Expression{}
	}

	e.init()
	return e
}

func (e *Expression) init() {
	defaultToEvery := func(field internal.Field) internal.Field {
		value, err := field.Format()
		if value == "" && err == nil {
			return internal.Every()
		}
		return field
	}

	e.Minute = defaultToEvery(e.Minute)
	e.Hour = defaultToEvery(e.Hour)
	e.Day = defaultToEvery(e.Day)
	e.Month = defaultToEvery(e.Month)
	e.Weekday = defaultToEvery(e.Weekday)
}

// Format Produce the crontab expression (eg "* * * * *")
// if any errors occurred during field generation (eg hour=35)
// then the first one encountered will be returned by this function.
func (e *Expression) Format() (string, error) {
	e.init()

	if e.Minute.Error() != nil {
		return "", e.Minute.Error()
	}
	if e.Hour.Error() != nil {
		return "", e.Hour.Error()
	}
	if e.Day.Error() != nil {
		return "", e.Day.Error()
	}
	if e.Month.Error() != nil {
		return "", e.Month.Error()
	}
	if e.Weekday.Error() != nil {
		return "", e.Weekday.Error()
	}

	return fmt.Sprintf(
		"%s %s %s %s %s",
		e.Minute.MustFormat(),
		e.Hour.MustFormat(),
		e.Day.MustFormat(),
		e.Month.MustFormat(),
		e.Weekday.MustFormat(),
	), nil
}

// MustFormat Similar to Format() but will panic if any invalid values were provided,
// ONLY use this method with known safe values, if you need to generate expressions
// from user input then use Format() and handle errors
func (e *Expression) MustFormat() string {
	result, err := e.Format()
	if err != nil {
		panic(fmt.Errorf("unhandled error, use Format() to handle: %w", err))
	}
	return result
}
