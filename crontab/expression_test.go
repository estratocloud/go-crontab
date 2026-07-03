package crontab

import (
	"errors"
	"reflect"
	"testing"
	"unsafe"

	"github.com/estratocloud/go-crontab/crontab/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getField(value string, err error) internal.Field {
	field := internal.Field{}
	v := reflect.ValueOf(&field).Elem()

	valuef := v.FieldByName("value")
	valueptr := unsafe.Pointer(valuef.UnsafeAddr())
	reflect.NewAt(valuef.Type(), valueptr).Elem().Set(reflect.ValueOf(value))

	if err != nil {
		errf := v.FieldByName("err")
		errptr := unsafe.Pointer(errf.UnsafeAddr())
		reflect.NewAt(errf.Type(), errptr).Elem().Set(reflect.ValueOf(err))
	}

	return field
}

func Test_NewExpression1(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(nil)
	assert.Equal(t, want, got)
}

func Test_NewExpression2(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{})
	assert.Equal(t, want, got)
}

func Test_NewExpression3(t *testing.T) {
	want := &Expression{
		Minute:  getField("0", nil),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Minute: getField("0", nil),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression4(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    getField("0", nil),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Hour: getField("0", nil),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression5(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     getField("1", nil),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Day: getField("1", nil),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression6(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   getField("1", nil),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Month: getField("1", nil),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression7(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: getField("1", nil),
	}
	got := NewExpression(&Expression{
		Weekday: getField("1", nil),
	})
	assert.Equal(t, want, got)
}

func Test_Format1(t *testing.T) {
	want := "* * * * *"
	got, err := NewExpression(nil).Format()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_Format2(t *testing.T) {
	want := "* * * * *"
	e := Expression{}
	got, err := e.Format()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_Format3(t *testing.T) {
	want := errors.New("bad minute")
	field := getField("", want)

	got, err := NewExpression(&Expression{Minute: field}).Format()
	assert.Equal(t, "", got)
	assert.ErrorIs(t, err, want)
}

func Test_Format4(t *testing.T) {
	want := errors.New("bad hour")
	field := getField("", want)

	got, err := NewExpression(&Expression{Hour: field}).Format()
	assert.Equal(t, "", got)
	assert.ErrorIs(t, err, want)
}

func Test_Format5(t *testing.T) {
	want := errors.New("bad day")
	field := getField("", want)

	got, err := NewExpression(&Expression{Day: field}).Format()
	assert.Equal(t, "", got)
	assert.ErrorIs(t, err, want)
}

func Test_Format6(t *testing.T) {
	want := errors.New("bad month")
	field := getField("", want)

	got, err := NewExpression(&Expression{Month: field}).Format()
	assert.Equal(t, "", got)
	assert.ErrorIs(t, err, want)
}

func Test_Format7(t *testing.T) {
	want := errors.New("bad weekday")
	field := getField("", want)

	got, err := NewExpression(&Expression{Weekday: field}).Format()
	assert.Equal(t, "", got)
	assert.ErrorIs(t, err, want)
}

func Test_MustFormat1(t *testing.T) {
	want := "* * * * *"
	got := NewExpression(nil).MustFormat()
	assert.Equal(t, want, got)
}

func Test_MustFormat2(t *testing.T) {
	want := "* * * * *"
	e := Expression{}
	got := e.MustFormat()
	assert.Equal(t, want, got)
}

func Test_MustFormat3(t *testing.T) {
	want := errors.New("bad minute")
	field := getField("", want)

	e := NewExpression(&Expression{Minute: field})

	assert.PanicsWithErrorf(t, "unhandled error, use Format() to handle: bad minute", func() {
		_ = e.MustFormat()
	}, "MustFormat() should panic")
}
