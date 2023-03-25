package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsLeapYear(t *testing.T) {
	got := IsLeapYear(2000)
	assert.True(t, got)

	got = IsLeapYear(2001)
	assert.False(t, got)
}

func TestRoundToNearestMinute(t *testing.T) {
	testCases := []struct {
		minutes      int
		seconds      int
		want_hours   int
		want_minutes int
		want_seconds int
	}{
		{2, 29, 11, 2, 0},
		{2, 31, 11, 3, 0},
		{59, 31, 12, 0, 0},
	}
	for _, tc := range testCases {

		date := time.Date(1990, time.August, 13, 11, tc.minutes, tc.seconds, 0, time.UTC)

		got := RoundToNearestMinute(date)

		assert.Equal(t, tc.want_hours, got.Hour())
		assert.Equal(t, tc.want_minutes, got.Minute())
		assert.Equal(t, tc.want_seconds, got.Second())
	}
}

func TestResolveTime(t *testing.T) {
	dc := &DateComponents{Year: 1990, Month: 2, Day: 2}

	got := ResolveTimeByDateComponents(dc)

	want := time.Date(1990, time.February, 2, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, want, got)
}
