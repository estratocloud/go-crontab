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

func Test_Every1(t *testing.T) {
	field := Every()
	got := field.MustFormat()
	assert.Equal(t, "*", got)
}

func Test_Value1(t *testing.T) {
	field := Value("value1", 1, 2, 1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_Value2(t *testing.T) {
	field := Value("value2", 2, 4, 1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid value2 (1) must be between 2 and 4")
}

func Test_Value3(t *testing.T) {
	field := Value("value3", 1, 2, 3)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid value3 (3) must be between 1 and 2")
}

func Test_List1(t *testing.T) {
	field := List("list1", 1, 2, 1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_List2(t *testing.T) {
	field := List("list2", 1, 9, 4, 7, 9)
	got := field.MustFormat()
	assert.Equal(t, "4,7,9", got)
}

func Test_List3(t *testing.T) {
	immutableList := []int{1, 3, 2}
	field := List("list3", 1, 3, immutableList...)
	got := field.MustFormat()
	assert.Equal(t, "1,2,3", got)
	assert.Equal(t, []int{1, 3, 2}, immutableList)
}

func Test_List4(t *testing.T) {
	field := List("list4", 1, 3, 1, 2, 3, 4)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid list4 (4) must be between 1 and 3")
}

func Test_List5(t *testing.T) {
	field := List("list5", 5, 6, 1, 2, 3, 4, 5, 6)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid list5 (1) must be between 5 and 6")
}

func Test_List6(t *testing.T) {
	field := List("list6", 5, 6)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "no list6 values were passed")
}

func Test_Range1(t *testing.T) {
	field := Range("range1", 1, 5, 1, 5)
	got := field.MustFormat()
	assert.Equal(t, "1-5", got)
}

func Test_Range2(t *testing.T) {
	field := Range("range2", 1, 9, 4, 7)
	got := field.MustFormat()
	assert.Equal(t, "4-7", got)
}

func Test_Range3(t *testing.T) {
	field := Range("range3", 1, 3, 0, 3)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid range3 (0) must be between 1 and 3")
}

func Test_Range4(t *testing.T) {
	field := Range("range4", 1, 3, 1, 4)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid range4 (4) must be between 1 and 3")
}

func Test_Range5(t *testing.T) {
	field := Range("range5", 1, 3, 3, 1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid range5 range, 'from' (3) must be less than 'to' (1)")
}

func Test_EveryX1(t *testing.T) {
	field := EveryX("everyx1", 4, 2)
	got := field.MustFormat()
	assert.Equal(t, "*/2", got)
}

func Test_EveryX2(t *testing.T) {
	field := EveryX("everyx2", 4, 1)
	got := field.MustFormat()
	assert.Equal(t, "*", got)
}

func Test_EveryX3(t *testing.T) {
	field := EveryX("everyx3", 2, 0)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid everyx3 (0) must be between 1 and 2")
}

func Test_EveryX4(t *testing.T) {
	field := EveryX("everyx4", 12, 13)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid everyx4 (13) must be between 1 and 12")
}
