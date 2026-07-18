package crontab

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EveryXMinutes1(t *testing.T) {
	got, err := EveryXMinutes(5).Format()
	require.NoError(t, err)
	assert.Equal(t, "*/5 * * * *", got)
}

func Test_EveryXMinutes2(t *testing.T) {
	got, err := EveryXMinutes(1).Format()
	require.NoError(t, err)
	assert.Equal(t, "* * * * *", got)
}

func Test_EveryXMinutes3(t *testing.T) {
	got, err := EveryXMinutes(0).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (0) must be between 1 and 30")
}

func Test_EveryXMinutes4(t *testing.T) {
	got, err := EveryXMinutes(31).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (31) must be between 1 and 30")
}

func Test_EveryXHours1(t *testing.T) {
	got, err := EveryXHours(6).Format()
	require.NoError(t, err)
	assert.Equal(t, "0 */6 * * *", got)
}

func Test_EveryXHours2(t *testing.T) {
	got, err := EveryXHours(1).Format()
	require.NoError(t, err)
	assert.Equal(t, "0 * * * *", got)
}

func Test_EveryXHours3(t *testing.T) {
	got, err := EveryXHours(0).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid hour (0) must be between 1 and 12")
}

func Test_EveryXHours4(t *testing.T) {
	got, err := EveryXHours(13).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid hour (13) must be between 1 and 12")
}

func Test_EveryHour1(t *testing.T) {
	got, err := EveryHour(0).Format()
	require.NoError(t, err)
	assert.Equal(t, "0 * * * *", got)
}

func Test_EveryHour2(t *testing.T) {
	got, err := EveryHour(59).Format()
	require.NoError(t, err)
	assert.Equal(t, "59 * * * *", got)
}

func Test_EveryHour3(t *testing.T) {
	got, err := EveryHour(-1).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (-1) must be between 0 and 59")
}

func Test_EveryHour4(t *testing.T) {
	got, err := EveryHour(60).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (60) must be between 0 and 59")
}

func Test_EveryDayAt1(t *testing.T) {
	got, err := EveryDayAt(13, 30).Format()
	require.NoError(t, err)
	assert.Equal(t, "30 13 * * *", got)
}

func Test_EveryDayAt2(t *testing.T) {
	got, err := EveryDayAt(-1, 30).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid hour (-1) must be between 0 and 23")
}

func Test_EveryDayAt3(t *testing.T) {
	got, err := EveryDayAt(24, 30).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid hour (24) must be between 0 and 23")
}

func Test_EveryDayAt4(t *testing.T) {
	got, err := EveryDayAt(12, -1).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (-1) must be between 0 and 59")
}

func Test_EveryDayAt5(t *testing.T) {
	got, err := EveryDayAt(12, 60).Format()
	assert.Empty(t, got)
	assert.EqualError(t, err, "invalid minute (60) must be between 0 and 59")
}
