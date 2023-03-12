package adhango

import (
	"math"
)

// Radians converts `degrees` to the corresponding radian value.
func Radians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// Degrees converts `radians` to the corresponding degree value.
func Degrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

// MeanSolarLongitude returns the geometric mean longitude of the sun in degrees, given `t`, the julian century.
func MeanSolarLongitude(t float64) float64 {
	// Equation from Astronomical Algorithms page 163
	term1 := 280.4664567
	term2 := 36000.76983 * t
	term3 := 0.0003032 * math.Pow(t, 2)
	L0 := term1 + term2 + term3
	return UnwindAngle(L0)
}

// MeanLunarLongitude returns the geometric mean longitude of the moon in degrees, given `t`, the julian century.
func MeanLunarLongitude(t float64) float64 {
	// Equation from Astronomical Algorithms page 144
	term1 := 218.3165
	term2 := 481267.8813 * t
	Lp := term1 + term2
	return UnwindAngle(Lp)
}

// AscendingLunarNodeLongitude returns the ascending lunar node longitude, given `t`, the julian century,
func AscendingLunarNodeLongitude(t float64) float64 {
	// Equation from Astronomical Algorithms page 144
	term1 := 125.04452
	term2 := 1934.136261 * t
	term3 := 0.0020708 * math.Pow(t, 2)
	term4 := math.Pow(t, 3) / 450000
	Ω := term1 - term2 + term3 + term4
	return UnwindAngle(Ω)
}

// ApparentSolarLongitude returns apparent longitude of the Sun, referred to as the true equinox of	the date, given `t`, the julian century, and `L0`, the mean longitude.
func ApparentSolarLongitude(t float64, L0 float64) float64 {
	longitude := L0 + SolarEquationOfTheCenter(t, MeanSolarAnomaly(t))
	Ω := 125.04 - (1934.136 * t)
	λ := longitude - 0.00569 - (0.00478 * math.Sin(Radians(Ω)))
	return UnwindAngle(λ)
}

// SolarEquationOfTheCenter returns the sun's equation of the center in degrees, given `t`, the julian century, and `m`, the mean anomaly.
func SolarEquationOfTheCenter(t float64, m float64) float64 {
	// Equation from Astronomical Algorithms page 164
	Mrad := Radians(m)
	term1 := (1.914602 - (0.004817 * t) - (0.000014 * math.Pow(t, 2))) * math.Sin(Mrad)
	term2 := (0.019993 - (0.000101 * t)) * math.Sin(2*Mrad)
	term3 := 0.000289 * math.Sin(3*Mrad)
	return term1 + term2 + term3
}

// MeanSolarAnomaly returns the sun's equation of the center in degrees, given `t`, the julian century, and `m`, the mean anomaly.
func MeanSolarAnomaly(t float64) float64 {
	// Equation from Astronomical Algorithms page 163
	term1 := 357.52911
	term2 := 35999.05029 * t
	term3 := 0.0001537 * math.Pow(t, 2)
	M := term1 + term2 - term3
	return UnwindAngle(M)
}

// MeanSiderealTime returns mean sidereal time, aka the hour angle of the vernal equinox, in degrees, given `t`, the julian century.
func MeanSiderealTime(t float64) float64 {
	// Equation from Astronomical Algorithms page 165
	JD := (t * 36525) + 2451545.0
	term1 := 280.46061837
	term2 := 360.98564736629 * (JD - 2451545)
	term3 := 0.000387933 * math.Pow(t, 2)
	term4 := math.Pow(t, 3) / 38710000
	θ := term1 + term2 + term3 - term4
	return UnwindAngle(θ)
}

// NutationInLongitude returns the nutation in longitude, given `L0`, the the solar longitude, `Lp`, the lunar longitude and `Ω`, the ascending node.
func NutationInLongitude(L0 float64, Lp float64, Ω float64) float64 {
	// Equation from Astronomical Algorithms page 144
	term1 := (-17.2 / 3600) * math.Sin(Radians(Ω))
	term2 := (1.32 / 3600) * math.Sin(2*Radians(L0))
	term3 := (0.23 / 3600) * math.Sin(2*Radians(Lp))
	term4 := (0.21 / 3600) * math.Sin(2*Radians(Ω))
	return term1 - term2 - term3 + term4
}

// NutationInObliquity returns the nutation in obliquity, given `L0`, the the solar longitude, `Lp`, the lunar longitude and `Ω`, the ascending node.
func NutationInObliquity(L0 float64, Lp float64, Ω float64) float64 {
	// Equation from Astronomical Algorithms page 144
	term1 := (9.2 / 3600) * math.Cos(Radians(Ω))
	term2 := (0.57 / 3600) * math.Cos(2*Radians(L0))
	term3 := (0.10 / 3600) * math.Cos(2*Radians(Lp))
	term4 := (0.09 / 3600) * math.Cos(2*Radians(Ω))
	return term1 + term2 + term3 - term4
}

// MeanObliquityOfTheEcliptic returns the mean obliquity of the ecliptic in degrees, given `t`, the julian century.
func MeanObliquityOfTheEcliptic(t float64) float64 {
	// Equation from Astronomical Algorithms page 147
	term1 := 23.439291
	term2 := 0.013004167 * t
	term3 := 0.0000001639 * math.Pow(t, 2)
	term4 := 0.0000005036 * math.Pow(t, 3)
	return term1 - term2 - term3 + term4
}

// ApparentObliquityOfTheEcliptic returns the mean obliquity of the ecliptic, corrected for calculating the apparent position of the sun, in degrees, given `t` (julian century) and `ε0` (mean obliquity of the ecliptic).
func ApparentObliquityOfTheEcliptic(t float64, ε0 float64) float64 {
	// Equation from Astronomical Algorithms page 165
	O := 125.04 - (1934.136 * t)
	return ε0 + (0.00256 * math.Cos(Radians(O)))
}

// AltitudeOfCelestialBody returns the altitude of the celestial body,  given `φ` (the observer latitude), `δ` (the declination) and `H` (the local hour angle).
func AltitudeOfCelestialBody(φ float64, δ float64, H float64) float64 {
	// Equation from Astronomical Algorithms page 93
	term1 := math.Sin(Radians(φ)) * math.Sin(Radians(δ))
	term2 := math.Cos(Radians(φ)) * math.Cos(Radians(δ)) * math.Cos(Radians(H))
	return Degrees(math.Asin(term1 + term2))
}

// ApproximateTransit returns the approximate transite given L (the longitude), Θ0 (the sidereal time), and α2 (the right ascension).
func ApproximateTransit(L float64, Θ0 float64, α2 float64) float64 {
	// Equation from page Astronomical Algorithms 102
	Lw := L * -1
	return NormalizeWithBound((α2+Lw-Θ0)/360, 1)
}

// CorrectedTransit returns the time at which the sun is at its highest point in the sky (in universal time) given m0 (approximate transit), L (the longitude), Θ0 (the sidereal time), α2 (the right ascension), α1 (the previous right ascension) and α3 (the next right ascension).
func CorrectedTransit(m0 float64, L float64, Θ0 float64, α2 float64, α1 float64, α3 float64) float64 {
	// Equation from page Astronomical Algorithms 102
	Lw := L * -1
	θ := UnwindAngle(Θ0 + (360.985647 * m0))
	α := UnwindAngle(InterpolateAngles(α2, α1, α3, m0))
	H := ClosestAngle(θ - Lw - α)
	Δm := H / -360
	return (m0 + Δm) * 24
}

// Interpolate returns the interpolated value given y2 (the value), y1 (the previous value), y3 (the next value) and n (the factor).
func Interpolate(y2 float64, y1 float64, y3 float64, n float64) float64 {
	// Equation from Astronomical Algorithms page 24
	a := y2 - y1
	b := y3 - y2
	c := b - a
	return y2 + ((n / 2) * (a + b + (n * c)))
}

// InterpolateAngles returns the interpolation of three angles, accounting for angle unwinding, given y2 (the value), y1 (previousValue), y3 (nextValue), and n (factor).
func InterpolateAngles(y2 float64, y1 float64, y3 float64, n float64) float64 {
	// Equation from Astronomical Algorithms page 24
	a := UnwindAngle(y2 - y1)
	b := UnwindAngle(y3 - y2)
	c := b - a
	return y2 + ((n / 2) * (a + b + (n * c)))
}

// CorrectedHourAngle returns the corrected hour angle given m0 (the approximate transit), h0 (the angle), coordinates (the coordinates), afterTransit(whether it's after transit), Θ0 (the sidereal time), α2 (the right ascension), α1 (the previous right ascension), α3 (the next right ascension), δ2 (the declination), δ1 (the previous declination), and δ3 (the next declination).
func CorrectedHourAngle(m0 float64, h0 float64, coordinates *Coordinates, afterTransit bool, Θ0 float64, α2 float64, α1 float64, α3 float64, δ2 float64, δ1 float64, δ3 float64) float64 {
	// Equation from page Astronomical Algorithms 102
	Lw := coordinates.Longitude * -1
	term1 := math.Sin(Radians(h0)) - (math.Sin(Radians(coordinates.Latitude)) * math.Sin(Radians(δ2)))
	term2 := math.Cos(Radians(coordinates.Latitude)) * math.Cos(Radians(δ2))
	H0 := Degrees(math.Acos(term1 / term2))
	m := m0 + (H0 / 360)
	if !afterTransit {
		m = m0 - (H0 / 360)
	}
	θ := UnwindAngle(Θ0 + (360.985647 * m))
	α := UnwindAngle(InterpolateAngles(α2, α1, α3, m))
	δ := Interpolate(δ2, δ1, δ3, m)
	H := θ - Lw - α
	h := AltitudeOfCelestialBody(coordinates.Latitude, δ, H)
	term3 := h - h0
	term4 := 360 * math.Cos(Radians(δ)) * math.Cos(Radians(coordinates.Latitude)) * math.Sin(Radians(H))
	Δm := term3 / term4

	return (m + Δm) * 24
}
