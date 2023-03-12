package adhango

import "time"

type DateComponents struct {
	Year  int
	Month int
	Day   int
}

// NewDateComponents creates a new DateComponents using the year, month and day in `t`.
func NewDateComponents(t time.Time) *DateComponents {
	return &DateComponents{Year: t.Year(), Month: int(t.Month() + 1), Day: t.Day()}
}
