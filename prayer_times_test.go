package adhango

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func addSeconds(t time.Time, offset int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+offset, t.Second(), t.Nanosecond(), t.Location())
}

func TestPrayerTimes(t *testing.T) {
	date := NewDateComponents(time.Date(2015, time.Month(7), 12, 0, 0, 0, 0, time.UTC))
	params := GetMethodParameters(NORTH_AMERICA)
	params.Madhab = HANAFI

	coords, err := NewCoordinates(35.7750, -78.6336)
	assert.Nil(t, err)
	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Fajr)    // "04:42 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Sunrise) // "06:08 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Dhuhr)   // "01:21 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Asr)     // "06:22 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Maghrib) // "08:32 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(7), 12, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Isha)    // "09:57 PM" NYC time
}

func TestOffsets(t *testing.T) {
	date := NewDateComponents(time.Date(2015, time.Month(12), 1, 0, 0, 0, 0, time.UTC))
	coords, err := NewCoordinates(35.7750, -78.6336)
	assert.Nil(t, err)

	params := GetMethodParameters(MUSLIM_WORLD_LEAGUE)

	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	err = prayerTimes.SetTimeZone("America/New_York")
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2015, time.Month(12), 1, 5, 35, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Fajr)   // "05:35 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 7, 6, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Sunrise) // "07:06 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 12, 5, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Dhuhr)  // "12:05 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 2, 42, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Asr)    // "02:42 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 5, 1, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Maghrib) // "05:01 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 6, 26, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Isha)   // "06:26 PM" NYC time

	params.Adjustments.FajrAdj = 10
	params.Adjustments.SunriseAdj = 10
	params.Adjustments.DhuhrAdj = 10
	params.Adjustments.AsrAdj = 10
	params.Adjustments.MaghribAdj = 10
	params.Adjustments.IshaAdj = 10

	prayerTimes, err = NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	err = prayerTimes.SetTimeZone("America/New_York")
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2015, time.Month(12), 1, 5, 45, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Fajr)    // "05:45 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 7, 16, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Sunrise) // "07:16 AM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 12, 15, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Dhuhr)  // "12:15 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 2, 52, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Asr)     // "02:52 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 5, 11, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Maghrib) // "05:11 PM" NYC time
	assert.Equal(t, time.Date(2015, time.Month(12), 1, 6, 36, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Isha)    // "06:36 PM" NYC time
}

func TestMoonsightingMethod(t *testing.T) {
	date := NewDateComponents(time.Date(2016, time.Month(1), 31, 0, 0, 0, 0, time.UTC))
	coords, err := NewCoordinates(35.7750, -78.6336)
	assert.Nil(t, err)

	prayerTimes, err := NewPrayerTimes(coords, date, GetMethodParameters(MOON_SIGHTING_COMMITTEE))
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2016, time.Month(1), 31, 5, 48, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Fajr)    // "05:48 AM" NYC time
	assert.Equal(t, time.Date(2016, time.Month(1), 31, 7, 16, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Sunrise) // "07:16 AM" NYC time
	assert.Equal(t, time.Date(2016, time.Month(1), 31, 12, 33, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Dhuhr)  // "12:33 PM" NYC time
	assert.Equal(t, time.Date(2016, time.Month(1), 31, 3, 20, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Asr)     // "03:20 PM" NYC time
	assert.Equal(t, time.Date(2016, time.Month(1), 31, 5, 43, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Maghrib) // "05:43 PM" NYC time
	assert.Equal(t, time.Date(2016, time.Month(1), 31, 7, 5, 0, 0, time.FixedZone("America/New_York", 0)), prayerTimes.Isha)     // "07:05 PM" NYC time
}

func TestMoonsightingMethodHighLat(t *testing.T) {
	// Values from http://www.moonsighting.com/pray.php
	date := NewDateComponents(time.Date(2016, time.Month(1), 1, 0, 0, 0, 0, time.UTC))
	params := GetMethodParameters(MOON_SIGHTING_COMMITTEE)
	params.Madhab = HANAFI
	coords, err := NewCoordinates(59.9094, 10.7349)
	assert.Nil(t, err)

	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	err = prayerTimes.SetTimeZone("Europe/Oslo")
	assert.Nil(t, err)

	assert.Equal(t, time.Date(2016, time.Month(1), 1, 7, 34, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Fajr)    // "07:34 AM" Oslo time
	assert.Equal(t, time.Date(2016, time.Month(1), 1, 9, 19, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Sunrise) // "09:19 AM" Oslo time
	assert.Equal(t, time.Date(2016, time.Month(1), 1, 12, 25, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Dhuhr)  // "12:25 PM" Oslo time
	assert.Equal(t, time.Date(2016, time.Month(1), 1, 1, 36, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Asr)     // "01:36 PM" Oslo time
	assert.Equal(t, time.Date(2016, time.Month(1), 1, 3, 26, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Maghrib) // "03:25 PM" Oslo time
	assert.Equal(t, time.Date(2016, time.Month(1), 1, 5, 2, 0, 0, time.FixedZone("Europe/Oslo", 0)), prayerTimes.Isha)     // "05:02 PM" Oslo time
}

func TestTimeForPrayer(t *testing.T) {
	date := NewDateComponents(time.Date(2016, time.Month(7), 1, 0, 0, 0, 0, time.UTC))
	params := GetMethodParameters(MUSLIM_WORLD_LEAGUE)
	params.Madhab = HANAFI
	params.HighLatitudeRule = TWILIGHT_ANGLE
	coords, err := NewCoordinates(59.9094, 10.7349)
	assert.Nil(t, err)

	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	assert.Equal(t, prayerTimes.Fajr, prayerTimes.TimeForPrayer(FAJR))
	assert.Equal(t, prayerTimes.Sunrise, prayerTimes.TimeForPrayer(SUNRISE))
	assert.Equal(t, prayerTimes.Dhuhr, prayerTimes.TimeForPrayer(DHUHR))
	assert.Equal(t, prayerTimes.Asr, prayerTimes.TimeForPrayer(ASR))
	assert.Equal(t, prayerTimes.Maghrib, prayerTimes.TimeForPrayer(MAGHRIB))
	assert.Equal(t, prayerTimes.Isha, prayerTimes.TimeForPrayer(ISHA))
	assert.Nil(t, prayerTimes.TimeForPrayer(NO_PRAYER))
}

func TestCurrentPrayer(t *testing.T) {
	date := NewDateComponents(time.Date(2015, time.Month(9), 1, 0, 0, 0, 0, time.UTC))
	params := GetMethodParameters(KARACHI)
	params.Madhab = HANAFI
	params.HighLatitudeRule = TWILIGHT_ANGLE
	coords, err := NewCoordinates(33.720817, 73.090032)
	assert.Nil(t, err)

	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	assert.Equal(t, NO_PRAYER, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Fajr, -1)))
	assert.Equal(t, FAJR, prayerTimes.CurrentPrayer(prayerTimes.Fajr))
	assert.Equal(t, FAJR, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Fajr, 1)))
	assert.Equal(t, SUNRISE, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Sunrise, 1)))
	assert.Equal(t, DHUHR, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Dhuhr, 1)))
	assert.Equal(t, ASR, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Asr, 1)))
	assert.Equal(t, MAGHRIB, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Maghrib, 1)))
	assert.Equal(t, ISHA, prayerTimes.CurrentPrayer(addSeconds(prayerTimes.Isha, 1)))
}

func TestNextPrayer(t *testing.T) {
	date := NewDateComponents(time.Date(2015, time.Month(9), 1, 0, 0, 0, 0, time.UTC))
	params := GetMethodParameters(KARACHI)
	params.Madhab = HANAFI
	params.HighLatitudeRule = TWILIGHT_ANGLE
	coords, err := NewCoordinates(33.720817, 73.090032)
	assert.Nil(t, err)

	prayerTimes, err := NewPrayerTimes(coords, date, params)
	assert.Nil(t, err)

	assert.Equal(t, FAJR, prayerTimes.NextPrayer(addSeconds(prayerTimes.Fajr, -1)))
	assert.Equal(t, SUNRISE, prayerTimes.NextPrayer(prayerTimes.Fajr))
	assert.Equal(t, SUNRISE, prayerTimes.NextPrayer(addSeconds(prayerTimes.Fajr, 1)))
	assert.Equal(t, DHUHR, prayerTimes.NextPrayer(addSeconds(prayerTimes.Sunrise, 1)))
	assert.Equal(t, ASR, prayerTimes.NextPrayer(addSeconds(prayerTimes.Dhuhr, 1)))
	assert.Equal(t, MAGHRIB, prayerTimes.NextPrayer(addSeconds(prayerTimes.Asr, 1)))
	assert.Equal(t, ISHA, prayerTimes.NextPrayer(addSeconds(prayerTimes.Maghrib, 1)))
	assert.Equal(t, NO_PRAYER, prayerTimes.NextPrayer(addSeconds(prayerTimes.Isha, 1)))
}

func TestDaysSinceSolstice(t *testing.T) {
	// For Northern Hemisphere start from December 21
	// (DYY=0 for December 21, and counting forward, DYY=11 for January 1 and so on).
	// For Southern Hemisphere start from June 21
	// (DYY=0 for June 21, and counting forward)

	testCases := []struct {
		year     int
		month    int
		day      int
		latitude float64
		expected int
	}{
		{2016, 1, 1, 1.0, 11},
		{2015, 12, 31, 1.0, 10},
		{2016, 12, 31, 1.0, 10},
		{2016, 12, 21, 1.0, 0},
		{2016, 12, 22, 1.0, 1},
		{2016, 3, 1, 1.0, 71},
		{2015, 3, 1, 1.0, 70},
		{2016, 12, 20, 1.0, 365},
		{2015, 12, 20, 1.0, 364},
		{2015, 6, 21, -1.0, 0},
		{2016, 6, 21, -1.0, 0},
		{2015, 6, 20, -1.0, 364},
		{2016, 6, 20, -1.0, 365},
	}
	for _, tc := range testCases {
		date := time.Date(tc.year, time.Month(tc.month), tc.day, 0, 0, 0, 0, time.UTC)
		dss := DaysSinceSolstice(date.YearDay(), tc.year, tc.latitude)
		assert.Equal(t, tc.expected, dss)
	}
}
