package minute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Every1(t *testing.T) {
	field := Every()
	got := field.MustFormat()
	assert.Equal(t, "*", got)
}

func Test_Minute1(t *testing.T) {
	field := Minute(1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_Minute2(t *testing.T) {
	field := Minute(-1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (-1) must be between 0 and 59")
}

func Test_Minute3(t *testing.T) {
	field := Minute(60)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (60) must be between 0 and 59")
}

func Test_List1(t *testing.T) {
	field := List(1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_List2(t *testing.T) {
	field := List(4, 7, 9)
	got := field.MustFormat()
	assert.Equal(t, "4,7,9", got)
}

func Test_List3(t *testing.T) {
	immutableList := []int{1, 3, 2}
	field := List(immutableList...)
	got := field.MustFormat()
	assert.Equal(t, "1,2,3", got)
	assert.Equal(t, []int{1, 3, 2}, immutableList)
}

func Test_List4(t *testing.T) {
	field := List(1, 9, 59, 60)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (60) must be between 0 and 59")
}

func Test_List5(t *testing.T) {
	field := List(-1, 0, 1)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (-1) must be between 0 and 59")
}

func Test_List6(t *testing.T) {
	field := List()
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "no minute values were passed")
}

func Test_Range1(t *testing.T) {
	field := Range(1, 5)
	got := field.MustFormat()
	assert.Equal(t, "1-5", got)
}

func Test_Range2(t *testing.T) {
	field := Range(20, 59)
	got := field.MustFormat()
	assert.Equal(t, "20-59", got)
}

func Test_Range3(t *testing.T) {
	field := Range(-1, 3)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (-1) must be between 0 and 59")
}

func Test_Range4(t *testing.T) {
	field := Range(1, 60)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (60) must be between 0 and 59")
}

func Test_Range5(t *testing.T) {
	field := Range(4, 2)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute range, 'from' (4) must be less than 'to' (2)")
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
	assert.ErrorContains(t, err, "invalid minute (0) must be between 1 and 30")
}

func Test_EveryX4(t *testing.T) {
	field := EveryX(31)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid minute (31) must be between 1 and 30")
}
