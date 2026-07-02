package crontab

import (
	"testing"
	"time"

	"github.com/estratocloud/go-crontab/crontab/internal"

	"github.com/estratocloud/go-crontab/crontab/day"
	"github.com/estratocloud/go-crontab/crontab/hour"
	"github.com/estratocloud/go-crontab/crontab/minute"
	"github.com/estratocloud/go-crontab/crontab/month"
	"github.com/estratocloud/go-crontab/crontab/weekday"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		Minute:  minute.Minute(0),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Minute: minute.Minute(0),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression4(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    hour.Hour(0),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Hour: hour.Hour(0),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression5(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     day.Day(1),
		Month:   internal.Every(),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Day: day.Day(1),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression6(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   month.Month(time.January),
		Weekday: internal.Every(),
	}
	got := NewExpression(&Expression{
		Month: month.Month(time.January),
	})
	assert.Equal(t, want, got)
}

func Test_NewExpression7(t *testing.T) {
	want := &Expression{
		Minute:  internal.Every(),
		Hour:    internal.Every(),
		Day:     internal.Every(),
		Month:   internal.Every(),
		Weekday: weekday.Weekday(time.Sunday),
	}
	got := NewExpression(&Expression{
		Weekday: weekday.Weekday(time.Sunday),
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
	field := minute.Minute(333)
	got, err := NewExpression(&Expression{Minute: field}).Format()
	assert.Equal(t, "", got)
	assert.Errorf(t, err, "invalid minute (333) must be between 0 and 59")
}

func Test_Format4(t *testing.T) {
	field := hour.Hour(40)
	got, err := NewExpression(&Expression{Hour: field}).Format()
	assert.Equal(t, "", got)
	assert.Errorf(t, err, "invalid hour (40) must be between 0 and 23")
}

func Test_Format5(t *testing.T) {
	field := day.Day(50)
	got, err := NewExpression(&Expression{Day: field}).Format()
	assert.Equal(t, "", got)
	assert.Errorf(t, err, "invalid day (50) must be between 0 and 31")
}

func Test_Format6(t *testing.T) {
	field := month.Month(13)
	got, err := NewExpression(&Expression{Month: field}).Format()
	assert.Equal(t, "", got)
	assert.Errorf(t, err, "invalid month (13) must be between 1 and 12")
}

func Test_Format7(t *testing.T) {
	field := weekday.Weekday(7)
	got, err := NewExpression(&Expression{Weekday: field}).Format()
	assert.Equal(t, "", got)
	assert.Errorf(t, err, "invalid weekday (7) must be between 0 and 6")
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
	field := minute.Minute(300)

	e := NewExpression(&Expression{Minute: field})

	assert.PanicsWithErrorf(t, "unhandled error, use Format() to handle: invalid minute (300) must be between 0 and 59", func() {
		_ = e.MustFormat()
	}, "MustFormat() should panic")
}
