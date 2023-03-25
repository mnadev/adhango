package util

import "math"

type SolarCoordinates struct {
	// The declination of the sun, the angle between the rays of the Sun and the plane of the Earth's equator, in degrees.
	Declination float64
	// Right ascension of the Sun, the angular distance on the celestial equator from the vernal equinox to the hour circle, in degrees.
	RightAscension float64
	// Apparent sidereal time, the hour angle of the vernal equinox, in degrees.
	ApparentSiderealTime float64
}

func NewSolarCoordinates(julian_day float64) *SolarCoordinates {

	T := GetJulianCentury(julian_day)
	L0 := MeanSolarLongitude(T)
	Lp := MeanLunarLongitude(T)
	Ω := AscendingLunarNodeLongitude(T)
	λ := Radians(ApparentSolarLongitude(T, L0))
	θ0 := MeanSiderealTime(T)
	ΔΨ := NutationInLongitude(L0, Lp, Ω)
	Δε := NutationInObliquity(L0, Lp, Ω)
	ε0 := MeanObliquityOfTheEcliptic(T)
	εapp := Radians(ApparentObliquityOfTheEcliptic(T, ε0))

	// Equation from Astronomical Algorithms page 165
	declination := Degrees(math.Asin(math.Sin(εapp) * math.Sin(λ)))

	// Equation from Astronomical Algorithms page 165
	rightAscension := UnwindAngle(Degrees(math.Atan2(math.Cos(εapp)*math.Sin(λ), math.Cos(λ))))

	// Equation from Astronomical Algorithms page 88
	apparentSiderealTime := θ0 + (((ΔΨ * 3600) * math.Cos(Radians(ε0+Δε))) / 3600)

	return &SolarCoordinates{
		Declination:          declination,
		RightAscension:       rightAscension,
		ApparentSiderealTime: apparentSiderealTime,
	}
}
