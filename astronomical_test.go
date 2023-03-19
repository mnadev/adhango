package adhango

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func timeString(when float64) string {
	components, err := NewTimeComponents(when)
	if err != nil {
		return ""
	}

	minutes := components.Minutes + int(math.Round(float64(components.Seconds)/60.0))
	return fmt.Sprintf("%d:%02d", components.Hours, minutes)
}

func TestAngleConversion(t *testing.T) {
	got := Degrees(math.Pi)
	want := 180.0
	assert.InDelta(t, want, got, 0.00001)

	got = Degrees(math.Pi / 2)
	want = 90.0
	assert.InDelta(t, want, got, 0.00001)
}

func TestSolarCoordinates(t *testing.T) {
	// Values from Astronomical Algorithms, page 165.
	jd := GetJulianDay(1992, 10, 13, 0)
	solar := NewSolarCoordinates(jd)

	T := GetJulianCentury(jd)
	L0 := MeanSolarLongitude(T)
	ε0 := MeanObliquityOfTheEcliptic(T)
	εapp := ApparentObliquityOfTheEcliptic(T, ε0)
	M := MeanSolarAnomaly(T)
	C := SolarEquationOfTheCenter(T, M)
	λ := ApparentSolarLongitude(T, L0)
	δ := solar.Declination
	α := UnwindAngle(solar.RightAscension)

	assert.InDelta(t, -0.072183436, T, 0.00000000001)
	assert.InDelta(t, 201.80720, L0, 0.00001)
	assert.InDelta(t, 23.44023, ε0, 0.00001)
	assert.InDelta(t, 23.43999, εapp, 0.00001)
	assert.InDelta(t, 278.99397, M, 0.00001)
	assert.InDelta(t, -1.89732, C, 0.00001)

	// Lower accuracy than desired.
	assert.InDelta(t, 199.90895, λ, 0.00002)
	assert.InDelta(t, -7.78507, δ, 0.00001)
	assert.InDelta(t, 198.38083, α, 0.00001)

	// Values from Astronomical Algorithms, page 88.

	jd = GetJulianDay(1987, 4, 10, 0)
	solar = NewSolarCoordinates(jd)
	T = GetJulianCentury(jd)

	θ0 := MeanSiderealTime(T)
	θapp := solar.ApparentSiderealTime
	Ω := AscendingLunarNodeLongitude(T)
	ε0 = MeanObliquityOfTheEcliptic(T)
	L0 = MeanSolarLongitude(T)
	Lp := MeanLunarLongitude(T)
	ΔΨ := NutationInLongitude(L0, Lp, Ω)
	Δε := NutationInObliquity(L0, Lp, Ω)
	ε := ε0 + Δε

	assert.InDelta(t, 197.693195, θ0, 0.000001)
	assert.InDelta(t, 197.6922295833, θapp, 0.0001)

	// Values from Astronomical Algorithms, page 148.
	assert.InDelta(t, 11.2531, Ω, 0.0001)
	assert.InDelta(t, -0.0010522, ΔΨ, 0.0001)
	assert.InDelta(t, 0.0026230556, Δε, 0.00001)
	assert.InDelta(t, 23.4409463889, ε0, 0.000001)
	assert.InDelta(t, 23.4435694444, ε, 0.00001)
}

func TestRightAscensionEdgeCase(t *testing.T) {
	coordinates, err := NewCoordinates(35+47.0/60.0, -78-39.0/60.0)
	assert.Nil(t, err)

	var previousTime *SolarTime
	var currTime *SolarTime
	for i := 0; i < 365; i++ {
		currTime = NewSolarTime(NewDateComponents(time.Date(2016, 1, 1+i, 0, 0, 0, 0, time.UTC)), coordinates)
		if i > 0 {
			// Transit from one day to another should not differ more than one minute.
			assert.InDelta(t, previousTime.Transit, currTime.Transit, 1.0/60.0)

			// Sunrise and sunset from one day to another should not differ more than two minutes.
			assert.InDelta(t, previousTime.Sunrise, currTime.Sunrise, 2.0/60.0)
			assert.InDelta(t, previousTime.Sunset, currTime.Sunset, 2.0/60.0)
		}
		previousTime = currTime
	}
}

func TestAltitudeOfCelestialBody(t *testing.T) {
	φ := 38 + (55 / 60.0) + (17.0 / 3600)
	δ := -6 - (43 / 60.0) - (11.61 / 3600)
	H := 64.352133
	h := AltitudeOfCelestialBody(φ, δ, H)

	want := 15.1249
	assert.InDelta(t, want, h, 0.0001)
}

func TestTransitAndHourAngle(t *testing.T) {
	// Calues from Astronomical Algorithms, page 103.
	longitude := -71.0833
	Θ := 177.74208
	α1 := 40.68021
	α2 := 41.73129
	α3 := 42.78204
	m0 := ApproximateTransit(longitude, Θ, α2)

	want := 0.81965
	assert.InDelta(t, want, m0, 0.00001)

	transit := CorrectedTransit(m0, longitude, Θ, α2, α1, α3) / 24
	want = 0.81980
	assert.InDelta(t, want, transit, 0.00001)

	δ1 := 18.04761
	δ2 := 18.44092
	δ3 := 18.82742
	coords, err := NewCoordinates(42.3333, longitude)
	assert.Nil(t, err)

	rise := CorrectedHourAngle(m0, -0.5667, coords, false, Θ, α2, α1, α3, δ2, δ1, δ3) / 24
	want = 0.51766
	assert.InDelta(t, want, rise, 0.00001)
}

func TestInterpolation(t *testing.T) {
	// Values from Astronomical Algorithms, page 25.
	interpolatedValue := Interpolate(0.877366, 0.884226, 0.870531, 4.35/24)
	assert.InDelta(t, 0.876125, interpolatedValue, 0.000001)

	i1 := Interpolate(1, -1, 3, 0.6)
	assert.InDelta(t, 2.2, i1, 0.000001)
}

func TestAngleInterpolation(t *testing.T) {
	i1 := InterpolateAngles(1, -1, 3, 0.6)
	assert.InDelta(t, 2.2, i1, 0.000001)

	i2 := InterpolateAngles(1, 359, 3, 0.6)
	assert.InDelta(t, 2.2, i2, 0.000001)
}

func TestLeapYear(t *testing.T) {

	testCases := []struct {
		year int
		want bool
	}{
		{2015, false},
		{2016, true},
		{1600, true},
		{2000, true},
		{2400, true},
		{1700, false},
		{1800, false},
		{1900, false},
		{2100, false},
		{2200, false},
		{2300, false},
		{2600, false},
	}
	for _, tc := range testCases {
		got := IsLeapYear(tc.year)
		assert.Equal(t, tc.want, got)
	}
}
