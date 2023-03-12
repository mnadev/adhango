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
	got, err := NewTimeComponents(100000000.321)

	if err != nil {
		t.Fatalf("error with good float64, got err = %v", err)
	}

	want := TimeComponents{
		Hours:   100000000,
		Minutes: 19,
		Seconds: 15,
	}
	if want != *got {
		t.Fatalf("wrong time, got = %+v, want = %+v", got, want)
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
