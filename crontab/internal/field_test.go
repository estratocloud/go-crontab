package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error1(t *testing.T) {
	want := errors.New("test error")
	field := &Field{
		err: want,
	}
	got := field.Error()
	assert.ErrorIs(t, got, want)
}

func Test_Error2(t *testing.T) {
	field := &Field{}
	got := field.Error()
	assert.Nil(t, got)
}

func Test_Format1(t *testing.T) {
	wantError := errors.New("test error")
	field := &Field{
		value: "ok",
		err:   wantError,
	}
	gotValue, gotError := field.Format()
	assert.Equal(t, "ok", gotValue)
	assert.ErrorIs(t, gotError, wantError)
}

func Test_Format2(t *testing.T) {
	field := &Field{}
	gotValue, gotError := field.Format()
	assert.Equal(t, "", gotValue)
	assert.Nil(t, gotError)
}

func Test_MustFormat1(t *testing.T) {
	field := &Field{
		value: "* * * * *",
	}
	got := field.MustFormat()
	assert.Equal(t, "* * * * *", got)
}

func Test_MustFormat2(t *testing.T) {
	field := &Field{
		err: errors.New("test error"),
	}
	assert.PanicsWithErrorf(t, "unhandled error, use Format() to handle: test error", func() {
		_ = field.MustFormat()
	}, "MustFormat() should panic")
}
