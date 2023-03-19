package adhango

import (
	"fmt"
	"math"
	"testing"
	"time"
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
	if math.Abs(got-want) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", got, want)
	}

	got = Degrees(math.Pi / 2)
	want = 90.0
	if math.Abs(got-want) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", got, want)
	}
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

	if math.Abs(T-(-0.072183436)) > 0.00000000001 {
		t.Errorf("error; got = %.2f want = %.2f", T, -0.072183436)
	}

	if math.Abs(L0-201.80720) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", L0, 201.80720)
	}

	if math.Abs(ε0-23.44023) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", ε0, 23.44023)
	}

	if math.Abs(εapp-23.43999) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", εapp, 23.43999)
	}

	if math.Abs(M-278.99397) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", M, 278.99397)
	}

	if math.Abs(C-(-1.89732)) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", C, -1.89732)
	}

	// Lower accuracy than desired.
	if math.Abs(λ-199.90895) > 0.00002 {
		t.Errorf("error; got = %.2f want = %.2f", λ, 199.90895)
	}

	if math.Abs(δ-(-7.78507)) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", δ, -7.78507)
	}

	if math.Abs(α-198.38083) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", α, 198.38083)
	}

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

	if math.Abs(θ0-197.693195) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", θ0, 197.693195)
	}

	if math.Abs(θapp-197.6922295833) > 0.0001 {
		t.Errorf("error; got = %.2f want = %.2f", θapp, 197.6922295833)
	}

	// Values from Astronomical Algorithms, page 148.
	if math.Abs(Ω-11.2531) > 0.0001 {
		t.Errorf("error; got = %.2f want = %.2f", Ω, 11.2531)
	}

	if math.Abs(ΔΨ-(-0.0010522)) > 0.0001 {
		t.Errorf("error; got = %.2f want = %.2f", ΔΨ, -0.0010522)
	}

	if math.Abs(Δε-0.0026230556) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", Δε, 0.0026230556)
	}

	if math.Abs(ε0-23.4409463889) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", ε0, 23.4409463889)
	}

	if math.Abs(ε-23.4435694444) > 0.00001 {
		t.Errorf("error; got = %.2f want = %.2f", ε, 23.4435694444)
	}
}

func TestRightAscensionEdgeCase(t *testing.T) {
	coordinates, err := NewCoordinates(35+47.0/60.0, -78-39.0/60.0)
	if err != nil {
		t.Errorf("got error %+v", err)
	}

	var previousTime *SolarTime
	var currTime *SolarTime
	for i := 0; i < 365; i++ {
		currTime = NewSolarTime(NewDateComponents(time.Date(2016, 1, 1+i, 0, 0, 0, 0, time.UTC)), coordinates)
		if i > 0 {
			// Transit from one day to another should not differ more than one minute.
			if math.Abs(currTime.Transit-previousTime.Transit) > 1.0/60.0 {
				t.Errorf("error; got difference = %.2f; wanted difference = %.2f", math.Abs(currTime.Transit-previousTime.Transit), 1.0/60.0)
			}

			// Sunrise and sunset from one day to another should not differ more than two minutes.
			if math.Abs(currTime.Sunrise-previousTime.Sunrise) > 2.0/60.0 {
				t.Errorf("error; got difference = %.2f; wanted difference = %.2f", math.Abs(currTime.Sunrise-previousTime.Sunrise), 2.0/60.0)
			}

			if math.Abs(currTime.Sunset-previousTime.Sunset) > 2.0/60.0 {
				t.Errorf("error; got difference = %.2f; wanted difference = %.2f", math.Abs(currTime.Sunset-previousTime.Sunset), 2.0/60.0)
			}
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
	if math.Abs(h-want) > 0.0001 {
		t.Fatalf("error with AltitudeOfCelestialBody, got %.4f want %.4f within range of 0.0001", h, want)
	}
}

func TestTransitAndHourAngle(t *testing.T) {
	// values from Astronomical Algorithms page 103
	longitude := -71.0833
	Θ := 177.74208
	α1 := 40.68021
	α2 := 41.73129
	α3 := 42.78204
	m0 := ApproximateTransit(longitude /* siderealTime */, Θ /* rightAscension */, α2)

	want := 0.81965
	if math.Abs(m0-want) > 0.00001 {
		t.Fatalf("error with ApproximateTransit, got %.2f want %.2f within range of 0.00001", m0, want)
	}

	transit := CorrectedTransit( /* approximateTransit */ m0, longitude /* siderealTime */, Θ /* rightAscension */, α2 /* previousRightAscension */, α1 /* nextRightAscension */, α3) / 24
	want = 0.81980
	if math.Abs(transit-want) > 0.00001 {
		t.Fatalf("error with CorrectedTransit, got %.2f want %.2f within range of 0.00001", transit, want)
	}

	δ1 := 18.04761
	δ2 := 18.44092
	δ3 := 18.82742
	coords, err := NewCoordinates( /* latitude */ 42.3333, longitude)
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	rise := CorrectedHourAngle( /* approximateTransit */ m0,
		/* angle */ -0.5667, coords,
		/* afterTransit */ false /* siderealTime */, Θ,
		/* rightAscension */ α2 /* previousRightAscension */, α1,
		/* nextRightAscension */ α3 /* declination */, δ2,
		/* previousDeclination */ δ1 /* nextDeclination */, δ3) / 24
	want = 0.51766
	if math.Abs(rise-want) > 0.00001 {
		t.Fatalf("error with CorrectedHourAngle, got %.2f want %.2f within range of 0.00001", rise, want)
	}
}

func TestInterpolation(t *testing.T) {
	// values from Astronomical Algorithms page 25
	interpolatedValue := Interpolate( /* value */ 0.877366,
		/* previousValue */ 0.884226 /* nextValue */, 0.870531 /* factor */, 4.35/24)
	if math.Abs(interpolatedValue-0.876125) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", interpolatedValue, 0.876125)
	}

	i1 := Interpolate(
		/* value */ 1 /* previousValue */, -1 /* nextValue */, 3 /* factor */, 0.6)
	if math.Abs(i1-2.2) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", i1, 2.2)
	}
}

func TestAngleInterpolation(t *testing.T) {
	i1 := InterpolateAngles( /* value */ 1 /* previousValue */, -1,
		/* nextValue */ 3 /* factor */, 0.6)
	if math.Abs(i1-2.2) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", i1, 2.2)
	}

	i2 := InterpolateAngles( /* value */ 1 /* previousValue */, 359,
		/* nextValue */ 3 /* factor */, 0.6)
	if math.Abs(i2-2.2) > 0.000001 {
		t.Errorf("error; got = %.2f want = %.2f", i2, 2.2)
	}
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
		if got != tc.want {
			t.Errorf("error; got = %t want = %t", got, tc.want)
		}
	}
}
