package data

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeComponentsInfiniteNumber(t *testing.T) {
	_, err := NewTimeComponents(math.Inf(1))
	assert.Error(t, err)
}

func TestNewTimeComponentsNan(t *testing.T) {
	_, err := NewTimeComponents(math.NaN())
	assert.Error(t, err)
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

		assert.Nil(t, err)
		assert.Equal(t, tc.want_hours, got.Hours)
		assert.Equal(t, tc.want_minutes, got.Minutes)
		assert.Equal(t, tc.want_seconds, got.Seconds)
	}
}

func TestTimeComponentsDateComponentsSuccess(t *testing.T) {
	tc, err := NewTimeComponents(10.2)

	assert.Nil(t, err)

	date := &DateComponents{
		Year:  1965,
		Month: int(time.February),
		Day:   10,
	}

	got := tc.DateComponents(date)
	want := time.Date(1965, time.February, 10, 10, 11, 59, 0, time.UTC)
	assert.Equal(t, want, got)
}
