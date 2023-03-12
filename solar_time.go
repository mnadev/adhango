package adhango

import "math"

type SolarTime struct {
	Transit float64
	Sunrise float64
	Sunset  float64

	Observer           *Coordinates
	Solar              *SolarCoordinates
	PrevSolar          *SolarCoordinates
	NextSolar          *SolarCoordinates
	ApproximateTransit float64
}

func NewSolarTime(d *DateComponents, c *Coordinates) *SolarTime {
	julianDate := GetJulianDay(d.Year, d.Month, d.Day, 0)

	prevSolar := NewSolarCoordinates(julianDate - 1)
	solar := NewSolarCoordinates(julianDate)
	nextSolar := NewSolarCoordinates(julianDate + 1)

	approximateTransit := ApproximateTransit(c.Longitude, solar.ApparentSiderealTime, solar.RightAscension)
	solarAltitude := -50.0 / 60.0

	transit := CorrectedTransit(approximateTransit, c.Longitude, solar.ApparentSiderealTime, solar.RightAscension, prevSolar.RightAscension, nextSolar.RightAscension)
	sunrise := CorrectedHourAngle(approximateTransit, solarAltitude, c, false, solar.ApparentSiderealTime, solar.RightAscension, prevSolar.RightAscension, nextSolar.RightAscension, solar.Declination, prevSolar.Declination, nextSolar.Declination)
	sunset := CorrectedHourAngle(approximateTransit, solarAltitude, c, true, solar.ApparentSiderealTime, solar.RightAscension, prevSolar.RightAscension, nextSolar.RightAscension, solar.Declination, prevSolar.Declination, nextSolar.Declination)

	return &SolarTime{
		Transit:            transit,
		Sunrise:            sunrise,
		Sunset:             sunset,
		Observer:           c,
		Solar:              solar,
		PrevSolar:          prevSolar,
		NextSolar:          nextSolar,
		ApproximateTransit: approximateTransit,
	}
}

func (s *SolarTime) HourAngle(angle float64, afterTransit bool) float64 {
	return CorrectedHourAngle(s.ApproximateTransit, angle, s.Observer, afterTransit, s.Solar.ApparentSiderealTime, s.Solar.RightAscension, s.PrevSolar.RightAscension, s.NextSolar.RightAscension, s.Solar.Declination, s.PrevSolar.Declination, s.NextSolar.Declination)
}

func (s *SolarTime) Afternoon(sl ShadowLength) float64 {
	// TODO (from Swift version) source shadow angle calculation
	tangent := math.Abs(s.Observer.Latitude - s.Solar.Declination)
	inverse := ShadowLengthToFloatMap[sl] + math.Tan(Radians(tangent))
	angle := Degrees(math.Atan(1.0 / inverse))

	return s.HourAngle(angle, true)
}
