package adhango

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolarTime(t *testing.T) {
	/*
	 * Comparison values generated from
	 * http://aa.usno.navy.mil/rstt/onedaytable?form=1&ID=AA&year=2015&month=7&day=12&state=NC&place=raleigh
	 */

	coordinates, err := NewCoordinates(35+47.0/60.0, -78-39.0/60.0)
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	solar := NewSolarTime(NewDateComponents(time.Date(2015, time.Month(7), 12, 0, 0, 0, 0, time.UTC)), coordinates)

	transit := solar.Transit
	sunrise := solar.Sunrise
	sunset := solar.Sunset
	twilightStart := solar.HourAngle(-6 /* afterTransit */, false)
	twilightEnd := solar.HourAngle(-6 /* afterTransit */, true)
	invalid := solar.HourAngle(-36 /* afterTransit */, true)

	got := timeString(twilightStart)
	want := "9:38"
	assert.Equal(t, want, got)

	got = timeString(sunrise)
	want = "10:08"
	assert.Equal(t, want, got)

	got = timeString(transit)
	want = "17:20"
	assert.Equal(t, want, got)

	got = timeString(sunset)
	want = "24:32"
	assert.Equal(t, want, got)

	got = timeString(twilightEnd)
	want = "25:02"
	assert.Equal(t, want, got)

	got = timeString(invalid)
	want = ""
	assert.Equal(t, want, got)
}

func TestCalendricalDate(t *testing.T) {
	// generated from http://aa.usno.navy.mil/data/docs/RS_OneYear.php for KUKUIHAELE, HAWAII
	coordinates, err := NewCoordinates( /* latitude */ 20+7.0/60.0 /* longitude */, -155.0-34.0/60.0)
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	day1solar := NewSolarTime(NewDateComponents(time.Date(2015, time.Month(4), 2, 0, 0, 0, 0, time.UTC)), coordinates)
	day2solar := NewSolarTime(NewDateComponents(time.Date(2015, time.Month(4), 3, 0, 0, 0, 0, time.UTC)), coordinates)

	day1 := day1solar.Sunrise
	day2 := day2solar.Sunrise

	gotStrDay1 := timeString(day1)
	wantStrDay1 := "16:15"
	assert.Equal(t, wantStrDay1, gotStrDay1)

	gotStrDay2 := timeString(day2)
	wantStrDay2 := "16:14"
	assert.Equal(t, wantStrDay2, gotStrDay2)
}
