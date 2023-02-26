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
	date := time.Date(1990, time.August, 13, 10, 40, 31, 0, time.UTC)

	got := RoundToNearestMinute(date)

	want := time.Date(1990, time.August, 13, 10, 41, 0, 0, time.UTC)

	if want != got {
		t.Fatalf("wrong time, got = %+v, want = %+v", got, want)
	}
}

func TestResolveTime(t *testing.T) {
	dc := DateComponents{Year: 1990, Month: 2, Day: 2}

	got := ResolveTimeByDateComponents(dc)

	want := time.Date(1990, time.February, 2, 0, 0, 0, 0, time.UTC)

	if want != got {
		t.Fatalf("wrong time, got = %+v, want = %+v", got, want)
	}
}
