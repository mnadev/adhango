package adhango

import (
	"math"
	"time"
)

type PrayerTimes struct {
	Fajr              time.Time
	Sunrise           time.Time
	Dhuhr             time.Time
	Asr               time.Time
	Maghrib           time.Time
	Isha              time.Time
	Coords            *Coordinates
	DateComponent     *DateComponents
	CalculationParams *CalculationParameters
}

func NewPrayerTimes(coords *Coordinates, date *DateComponents, params *CalculationParameters) (*PrayerTimes, error) {
	prayerDate := ResolveTimeByDateComponents(date)
	dayOfYear := prayerDate.YearDay()

	tomorrowDate := prayerDate.AddDate(0, 0, 1)
	tomorrow := NewDateComponents(tomorrowDate)

	solarTime := NewSolarTime(date, coords)

	timeComponents, err := NewTimeComponents(solarTime.Transit)
	if err != nil {
		return nil, err
	}
	transit := timeComponents.DateComponents(date)

	timeComponents, err = NewTimeComponents(solarTime.Sunrise)
	if err != nil {
		return nil, err
	}
	sunriseComponents := timeComponents.DateComponents(date)

	timeComponents, err = NewTimeComponents(solarTime.Sunset)
	if err != nil {
		return nil, err
	}
	sunsetComponents := timeComponents.DateComponents(date)

	tomorrowSolarTime := NewSolarTime(tomorrow, coords)
	tomorrowSunriseComponents, err := NewTimeComponents(tomorrowSolarTime.Sunrise)
	if err != nil {
		return nil, err
	}

	tempDhuhr := transit
	tempSunrise := sunriseComponents
	tempMaghrib := sunsetComponents

	timeComponents, err = NewTimeComponents(solarTime.Afternoon(MadhabToShadowLengthMap[params.Madhab]))
	if err != nil {
		return nil, err
	}
	tempAsr := timeComponents.DateComponents(date)

	tomorrowSunrise := tomorrowSunriseComponents.DateComponents(tomorrow)
	night := tomorrowSunrise.Sub(sunsetComponents)

	tempFajr := time.Time{}
	timeComponents, err = NewTimeComponents(solarTime.HourAngle(-1*params.FajrAngle, false))
	if err == nil {
		tempFajr = timeComponents.DateComponents(date)
	}

	if params.Method == MOON_SIGHTING_COMMITTEE && coords.Latitude >= 55 {
		tempFajr = sunriseComponents.Add(time.Second * time.Duration(-1*(int)(night/7000)))
	}

	nightPortions, err := params.NightPortions()
	if err != nil {
		return nil, err
	}

	safeFajr := time.Time{}
	if params.Method == MOON_SIGHTING_COMMITTEE {
		safeFajr = SeasonAdjustedMorningTwilight(coords.Latitude, dayOfYear, date.Year, sunriseComponents)
	} else {
		portion := nightPortions.Fajr
		nightFraction := (int64)(portion * night.Seconds() / 1000)
		safeFajr = sunriseComponents.Add(time.Second * time.Duration(-1*(int)(nightFraction)))
	}

	if tempFajr.IsZero() || tempFajr.Before(safeFajr) {
		tempFajr = safeFajr
	}

	// Isha calculation with check against safe value
	var tempIsha time.Time
	if params.IshaInterval > 0 {
		tempIsha = tempMaghrib.Add(time.Second * time.Duration(params.IshaInterval*60))
	} else {
		tempIsha := time.Time{}
		timeComponents, err = NewTimeComponents(solarTime.HourAngle(-1*params.IshaAngle, true))
		if err == nil {
			tempIsha = timeComponents.DateComponents(date)
		}

		if params.Method == MOON_SIGHTING_COMMITTEE && coords.Latitude >= 55 {
			nightFraction := int64(night / 7000)
			tempIsha = sunsetComponents.Add(time.Second * time.Duration(nightFraction))
		}

		safeIsha := time.Time{}
		if params.Method == MOON_SIGHTING_COMMITTEE {
			safeIsha = SeasonAdjustedEveningTwilight(coords.Latitude, dayOfYear, date.Year, sunsetComponents)
		} else {
			portion := nightPortions.Isha
			nightFraction := int64(portion * night.Seconds() / 1000)
			safeIsha = sunsetComponents.Add(time.Second * time.Duration(nightFraction))
		}

		if tempIsha.IsZero() || tempIsha.After(safeIsha) {
			tempIsha = safeIsha
		}
	}

	// Assign final times to public struct members with all offsets
	fajr := RoundToNearestMinute(tempFajr.Add(time.Minute * time.Duration(params.Adjustments.FajrAdj+params.MethodAdjustments.FajrAdj)))
	sunrise := RoundToNearestMinute(tempSunrise.Add(time.Minute * time.Duration(params.Adjustments.SunriseAdj+params.MethodAdjustments.SunriseAdj)))
	dhuhr := RoundToNearestMinute(tempDhuhr.Add(time.Minute * time.Duration(params.Adjustments.DhuhrAdj+params.MethodAdjustments.DhuhrAdj)))
	asr := RoundToNearestMinute(tempAsr.Add(time.Minute * time.Duration(params.Adjustments.AsrAdj+params.MethodAdjustments.AsrAdj)))
	maghrib := RoundToNearestMinute(tempMaghrib.Add(time.Minute * time.Duration(params.Adjustments.MaghribAdj+params.MethodAdjustments.MaghribAdj)))
	isha := RoundToNearestMinute(tempIsha.Add(time.Minute * time.Duration(params.Adjustments.IshaAdj+params.MethodAdjustments.IshaAdj)))

	return &PrayerTimes{
		Fajr:              fajr,
		Sunrise:           sunrise,
		Dhuhr:             dhuhr,
		Asr:               asr,
		Maghrib:           maghrib,
		Isha:              isha,
		Coords:            coords,
		DateComponent:     date,
		CalculationParams: params,
	}, nil
}

func (p *PrayerTimes) CurrentPrayerNow() Prayer {
	return p.CurrentPrayer(time.Now().UTC())
}

func (p *PrayerTimes) CurrentPrayer(t time.Time) Prayer {
	if p.Isha.Unix()-t.Unix() <= 0 {
		return ISHA
	} else if p.Maghrib.Unix()-t.Unix() <= 0 {
		return MAGHRIB
	} else if p.Asr.Unix()-t.Unix() <= 0 {
		return ASR
	} else if p.Dhuhr.Unix()-t.Unix() <= 0 {
		return DHUHR
	} else if p.Sunrise.Unix()-t.Unix() <= 0 {
		return SUNRISE
	} else if p.Fajr.Unix()-t.Unix() <= 0 {
		return FAJR
	} else {
		return NO_PRAYER
	}
}

func (p *PrayerTimes) NextPrayerNow() Prayer {
	return p.NextPrayer(time.Now().UTC())
}

func (p *PrayerTimes) NextPrayer(t time.Time) Prayer {
	if p.Isha.Unix()-t.Unix() <= 0 {
		return NO_PRAYER
	} else if p.Maghrib.Unix()-t.Unix() <= 0 {
		return ISHA
	} else if p.Asr.Unix()-t.Unix() <= 0 {
		return MAGHRIB
	} else if p.Dhuhr.Unix()-t.Unix() <= 0 {
		return ASR
	} else if p.Sunrise.Unix()-t.Unix() <= 0 {
		return DHUHR
	} else if p.Fajr.Unix()-t.Unix() <= 0 {
		return SUNRISE
	} else {
		return FAJR
	}
}

func (p *PrayerTimes) TimeForPrayer(prayer Prayer) time.Time {
	switch prayer {
	case FAJR:
		return p.Fajr
	case SUNRISE:
		return p.Sunrise
	case DHUHR:
		return p.Dhuhr
	case ASR:
		return p.Asr
	case MAGHRIB:
		return p.Maghrib
	case ISHA:
		return p.Isha
	case NO_PRAYER:
	default:
		break
	}
	return time.Time{}
}

func SeasonAdjustedMorningTwilight(latitude float64, day int, year int, sunrise time.Time) time.Time {
	a := 75.0 + ((28.65 / 55.0) * math.Abs(latitude))
	b := 75.0 + ((19.44 / 55.0) * math.Abs(latitude))
	c := 75.0 + ((32.74 / 55.0) * math.Abs(latitude))
	d := 75.0 + ((48.10 / 55.0) * math.Abs(latitude))

	adjustment := 0.0
	dyy := DaysSinceSolstice(day, year, latitude)
	if dyy < 91 {
		adjustment = a + (b-a)/91.0*float64(dyy)
	} else if dyy < 137 {
		adjustment = b + (c-b)/46.0*(float64(dyy)-91.0)
	} else if dyy < 183 {
		adjustment = c + (d-c)/46.0*(float64(dyy)-137.0)
	} else if dyy < 229 {
		adjustment = d + (c-d)/46.0*(float64(dyy)-183.0)
	} else if dyy < 275 {
		adjustment = c + (b-c)/46.0*(float64(dyy)-229.0)
	} else {
		adjustment = b + (a-b)/91.0*(float64(dyy)-275.0)
	}

	return sunrise.Add(time.Second * time.Duration(-1*math.Round(adjustment*60.0)))
}

func SeasonAdjustedEveningTwilight(latitude float64, day int, year int, sunset time.Time) time.Time {
	a := 75 + ((25.60 / 55.0) * math.Abs(latitude))
	b := 75 + ((2.050 / 55.0) * math.Abs(latitude))
	c := 75 - ((9.210 / 55.0) * math.Abs(latitude))
	d := 75 + ((6.140 / 55.0) * math.Abs(latitude))

	adjustment := 0.0
	dyy := DaysSinceSolstice(day, year, latitude)
	if dyy < 91 {
		adjustment = a + (b-a)/91.0*float64(dyy)
	} else if dyy < 137 {
		adjustment = b + (c-b)/46.0*(float64(dyy)-91.0)
	} else if dyy < 183 {
		adjustment = c + (d-c)/46.0*(float64(dyy)-137.0)
	} else if dyy < 229 {
		adjustment = d + (c-d)/46.0*(float64(dyy)-183.0)
	} else if dyy < 275 {
		adjustment = c + (b-c)/46.0*(float64(dyy)-229.0)
	} else {
		adjustment = b + (a-b)/91.0*(float64(dyy)-275.0)
	}

	return sunset.Add(time.Second * time.Duration(math.Round(adjustment*60.0)))
}

func DaysSinceSolstice(dayOfYear int, year int, latitude float64) int {
	daysSinceSolistice := 0
	northernOffset := 10
	isLeapYear := IsLeapYear(year)
	southernOffset := 172
	daysInYear := 365
	if isLeapYear {
		southernOffset = 173
		daysInYear = 366
	}

	if latitude >= 0 {
		daysSinceSolistice = dayOfYear + northernOffset
		if daysSinceSolistice >= daysInYear {
			daysSinceSolistice = daysSinceSolistice - daysInYear
		}
	} else {
		daysSinceSolistice = dayOfYear - southernOffset
		if daysSinceSolistice < 0 {
			daysSinceSolistice = daysSinceSolistice + daysInYear
		}
	}
	return daysSinceSolistice
}
