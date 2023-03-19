package adhango

import (
	"math"
	"testing"
)

func TestNightPortion(t *testing.T) {
	parameters := NewCalculationParametersBuilder().
		SetFajrAngle(18.0).
		SetIshaAngle(18.0).Build()
	parameters.HighLatitudeRule = MIDDLE_OF_THE_NIGHT
	np, err := parameters.NightPortions()
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	if math.Abs(np.Fajr-0.5) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Fajr, 0.5)
	}
	if math.Abs(np.Isha-0.5) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Isha, 0.5)
	}

	parameters = NewCalculationParametersBuilder().
		SetFajrAngle(18.0).
		SetIshaAngle(18.0).Build()
	parameters.HighLatitudeRule = SEVENTH_OF_THE_NIGHT
	np, err = parameters.NightPortions()
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	if math.Abs(np.Fajr-1.0/7.0) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Fajr, 1.0/7.0)
	}
	if math.Abs(np.Isha-1.0/7.0) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Isha, 1.0/7.0)
	}

	parameters = NewCalculationParametersBuilder().
		SetFajrAngle(10.0).
		SetIshaAngle(15.0).Build()
	parameters.HighLatitudeRule = TWILIGHT_ANGLE

	np, err = parameters.NightPortions()
	if err != nil {
		t.Errorf("got error %+v", err)
	}
	if math.Abs(np.Fajr-10.0/60.0) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Fajr, 10.0/60.0)
	}
	if math.Abs(np.Isha-15.0/60.0) > 0.001 {
		t.Errorf("error; got = %.2f wanted = %.2f", np.Isha, 15.0/60.0)
	}
}
