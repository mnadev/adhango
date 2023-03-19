package adhango

import (
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	got := IsLeapYear(2000)

	if !got {
		t.Fatalf("error with IsLeapYear, expected true")
	}

	got = IsLeapYear(2001)

	if got {
		t.Fatalf("error with IsLeapYear, expected false")
	}
}

func TestRoundToNearestMinute(t *testing.T) {
	testCases := []struct {
		minutes      int
		seconds      int
		want_hours   int
		want_minutes int
		want_seconds int
	}{
		{59, 31, 11, 2, 0},
		{2, 31, 11, 3, 0},
		{59, 31, 10, 59, 0},
	}
	for _, tc := range testCases {

		date := time.Date(1990, time.August, 13, 11, tc.minutes, tc.seconds, 0, time.UTC)

		got := RoundToNearestMinute(date)

		if tc.want_hours != got.Hour() {
			t.Fatalf("wrong hours; got = %d, want = %d %d", got.Hour(), tc.want_hours, date.Hour())
		}
		if tc.want_minutes != got.Minute() {
			t.Fatalf("wrong minute; got = %d, want = %d", got.Minute(), tc.want_minutes)
		}
		if tc.want_seconds != got.Second() {
			t.Fatalf("wrong seconds; got = %d, want = %d", got.Second(), tc.want_seconds)
		}
	}
}

func TestResolveTime(t *testing.T) {
	dc := &DateComponents{Year: 1990, Month: 2, Day: 2}

	got := ResolveTimeByDateComponents(dc)

	want := time.Date(1990, time.January, 2, 0, 0, 0, 0, time.UTC)

	if want != got {
		t.Fatalf("wrong time, got = %+v, want = %+v", got, want)
	}
}
