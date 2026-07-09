package month

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

func Test_Month1(t *testing.T) {
	field := Month(1)
	got := field.MustFormat()
	assert.Equal(t, "1", got)
}

func Test_Month2(t *testing.T) {
	field := Month(0)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (0) must be between 1 and 12")
}

func Test_Month3(t *testing.T) {
	field := Month(13)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (13) must be between 1 and 12")
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
	immutableList := []time.Month{1, 3, 2}
	field := List(immutableList...)
	got := field.MustFormat()
	assert.Equal(t, "1,2,3", got)
	assert.Equal(t, []time.Month{1, 3, 2}, immutableList)
}

func Test_List4(t *testing.T) {
	field := List(1, 9, 12, 13)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (13) must be between 1 and 12")
}

func Test_List5(t *testing.T) {
	field := List(0, 1, 2)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (0) must be between 1 and 12")
}

func Test_List6(t *testing.T) {
	field := List()
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "no month values were passed")
}

func Test_Range1(t *testing.T) {
	field := Range(1, 5)
	got := field.MustFormat()
	assert.Equal(t, "1-5", got)
}

func Test_Range2(t *testing.T) {
	field := Range(4, 8)
	got := field.MustFormat()
	assert.Equal(t, "4-8", got)
}

func Test_Range3(t *testing.T) {
	field := Range(0, 3)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (0) must be between 1 and 12")
}

func Test_Range4(t *testing.T) {
	field := Range(1, 13)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (13) must be between 1 and 12")
}

func Test_Range5(t *testing.T) {
	field := Range(4, 2)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month range, 'from' (4) must be less than 'to' (2)")
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
	assert.ErrorContains(t, err, "invalid month (0) must be between 1 and 6")
}

func Test_EveryX4(t *testing.T) {
	field := EveryX(7)
	got, err := field.Format()
	assert.Equal(t, "", got)
	assert.ErrorContains(t, err, "invalid month (7) must be between 1 and 6")
}
