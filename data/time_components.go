package adhango

import (
	"fmt"
	"math"
	"time"
)

type TimeComponents struct {
	Hours   int
	Minutes int
	Seconds int
}

// NewTimeComponents creates a TimeComponents struct given `d`, representing hours.
// An error is returned if `d` is infinite or NaN.
func NewTimeComponents(d float64) (*TimeComponents, error) {
	if math.IsInf(d, 0) {
		return nil, fmt.Errorf("given float64 value is infinite")
	}
	if math.IsNaN(d) {
		return nil, fmt.Errorf("given float64 value is NaN")
	}

	hours := math.Floor(d)
	minutes := math.Floor((d - hours) * 60.0)
	seconds := math.Floor((d - (hours + minutes/60.0)) * 60.0 * 60.0)
	return &TimeComponents{Hours: int(hours), Minutes: int(minutes), Seconds: int(seconds)}, nil
}

// DateComponents takes `d` and `t` and returns a time.Time struct using the values in the two given structs.
func (t *TimeComponents) DateComponents(d DateComponents) time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, t.Hours, t.Minutes, t.Seconds, 0, time.UTC)
}
