package adhango

import (
	"math"
	"testing"
	"time"
)

func TestNewTimeComponentsInfiniteNumber(t *testing.T) {
	_, err := NewTimeComponents(math.Inf(1))

	if err == nil {
		t.Fatalf("error with inf val, got nil")
	}
}

func TestNewTimeComponentsNan(t *testing.T) {
	_, err := NewTimeComponents(math.NaN())

	if err == nil {
		t.Fatalf("error with NaN val, got nil")
	}
}

func TestNewTimeComponentsSuccess(t *testing.T) {
	testCases := []struct {
		value        float64
		want_hours   int
		want_minutes int
		want_seconds int
	}{
		{15.199, 15, 11, 56},
		{1.0084, 1, 0, 30},
		{1.0083, 1, 0, 29},
		{2.1, 2, 6, 0},
		{3.5, 3, 30, 0},
	}
	for _, tc := range testCases {
		got, err := NewTimeComponents(tc.value)

		if err != nil {
			t.Fatalf("got err = %v", err)
		}

		if tc.want_hours != got.Hours {
			t.Fatalf("wrong hours; got = %d, want = %d", got.Hours, tc.want_hours)
		}
		if tc.want_minutes != got.Minutes {
			t.Fatalf("wrong minute; got = %d, want = %d", got.Minutes, tc.want_minutes)
		}
		if tc.want_seconds != got.Seconds {
			t.Fatalf("wrong seconds; got = %d, want = %d", got.Seconds, tc.want_seconds)
		}
	}
}

func TestTimeComponentsDateComponentsSuccess(t *testing.T) {
	tc, err := NewTimeComponents(10.2)

	if err != nil {
		t.Fatalf("error with good float64, got err = %v", err)
	}

	date := &DateComponents{
		Year:  1965,
		Month: int(time.February),
		Day:   10,
	}

	got := tc.DateComponents(date)

	want := time.Date(1965, time.February, 10, 10, 11, 59, 0, time.UTC)

	if want != got {
		t.Fatalf("wrong time, got = %+v, want = %+v", got, want)
	}
}
