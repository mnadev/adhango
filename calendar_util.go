package adhango

import (
	"math"
	"time"
)

// IsLeapYear returns true if `year` is a leap year, and false otherwise.
func IsLeapYear(year int) bool {
	return year%4 == 0 && !(year%100 == 0 && year%400 != 0)
}

// RoundToNearestMinute rounds to the closest minute for `d`.
func RoundToNearestMinute(d time.Time) time.Time {
	nearestMinute := d.Minute() + int(math.Round(float64(d.Second())/float64(60)))
	new_d := time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), nearestMinute, 0, 0, d.Location())
	return new_d
}

// ResolveTimeByDateComponents converts `c` to the corresponding time.Time value.
func ResolveTimeByDateComponents(c *DateComponents) time.Time {
	return ResolveTime(c.Year, c.Month, c.Day)
}

// ResolveTime creates a time.Time struct given `year`, `month` and `day`.
func ResolveTime(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
