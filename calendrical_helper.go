package adhango

import (
	"math"
)

// GetJulianDay returns the Julian day for a given Gregorian date.
func GetJulianDay(year int, month int, day int, hours float64) float64 {
	y := year
	if month <= 2 {
		y = year - 1
	}

	m := month
	if month <= 2 {
		m = month + 12
	}

	d := float64(day) + (hours / 24)

	a := math.Floor(float64(y) / 100.0)
	b := math.Floor(2 - a + (a / 4))

	i0 := int(365.25 * float64(y+4716))
	i1 := int(30.6001 * float64(m+1))

	return float64(i0+i1) + b + d - 1524.5
}

// GetJulianCentury returns the Julian century from epoch, given `julianDay`.
func GetJulianCentury(julianDay float64) float64 {
	// Equation from Astronomical Algorithms page 163
	return (julianDay - 2451545.0) / 36525
}
