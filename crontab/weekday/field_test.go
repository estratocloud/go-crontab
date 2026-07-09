package weekday

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Every1(t *testing.T) {
	field := Every()
	got := field.MustFormat()
	assert.Equal(t, "*", got)
}

func Test_Weekday1(t *testing.T) {
	field := Weekday(1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_Weekday2(t *testing.T) {
	field := Weekday(-1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (-1) must be between 0 and 6")
}

func Test_Weekday3(t *testing.T) {
	field := Weekday(7)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (7) must be between 0 and 6")
}

func Test_List1(t *testing.T) {
	field := List(1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_List2(t *testing.T) {
	field := List(0, 3, 6)
	got := field.MustFormat()
	assert.Equal(t, "0,3,6", got)
}

func Test_List3(t *testing.T) {
	immutableList := []time.Weekday{1, 3, 2}
	field := List(immutableList...)
	got := field.MustFormat()
	assert.Equal(t, "1,2,3", got)
	assert.Equal(t, []time.Weekday{1, 3, 2}, immutableList)
}

func Test_List4(t *testing.T) {
	field := List(1, 4, 6, 7)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (7) must be between 0 and 6")
}

func Test_List5(t *testing.T) {
	field := List(-1, 0, 1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (-1) must be between 0 and 6")
}

func Test_List6(t *testing.T) {
	field := List()
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "no weekday values were passed")
}

func Test_Range1(t *testing.T) {
	field := Range(1, 5)
	got := field.MustFormat()
	assert.Equal(t, "1-5", got)
}

func Test_Range2(t *testing.T) {
	field := Range(3, 6)
	got := field.MustFormat()
	assert.Equal(t, "3-6", got)
}

func Test_Range3(t *testing.T) {
	field := Range(-1, 3)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (-1) must be between 0 and 6")
}

func Test_Range4(t *testing.T) {
	field := Range(1, 7)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (7) must be between 0 and 6")
}

func Test_Range5(t *testing.T) {
	field := Range(4, 2)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday range, 'from' (4) must be less than 'to' (2)")
}

func Test_EveryX1(t *testing.T) {
	field := EveryX(2)
	got := field.MustFormat()
	assert.Equal(t, "*/2", got)
}

func Test_EveryX2(t *testing.T) {
	field := EveryX(1)
	got := field.MustFormat()
	assert.Equal(t, "*", got)
}

func Test_EveryX3(t *testing.T) {
	field := EveryX(0)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (0) must be between 1 and 3")
}

func Test_EveryX4(t *testing.T) {
	field := EveryX(4)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid weekday (4) must be between 1 and 3")
}
