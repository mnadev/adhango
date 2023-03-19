package adhango

import (
	"math"
	"testing"
)

func TestCalculationMethods(t *testing.T) {
	params := GetMethodParameters(MUSLIM_WORLD_LEAGUE)
	if math.Abs(params.FajrAngle-18) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.0)
	}
	if math.Abs(params.IshaAngle-17) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 17.0)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != MUSLIM_WORLD_LEAGUE {
		t.Errorf("error; got method = %d, want method %d", params.Method, MUSLIM_WORLD_LEAGUE)
	}

	params = GetMethodParameters(EGYPTIAN)
	if math.Abs(params.FajrAngle-19.5) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 19.5)
	}
	if math.Abs(params.IshaAngle-17.5) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 17.5)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != EGYPTIAN {
		t.Errorf("error; got method = %d, want method %d", params.Method, EGYPTIAN)
	}

	params = GetMethodParameters(KARACHI)
	if math.Abs(params.FajrAngle-18) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.0)
	}
	if math.Abs(params.IshaAngle-18) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 18.0)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != KARACHI {
		t.Errorf("error; got method = %d, want method %d", params.Method, KARACHI)
	}

	params = GetMethodParameters(UMM_AL_QURA)
	if math.Abs(params.FajrAngle-18.5) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.5)
	}
	if math.Abs(params.IshaAngle-0) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 0.0)
	}
	if float64(params.IshaInterval) != 90 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 90)
	}
	if params.Method != UMM_AL_QURA {
		t.Errorf("error; got method = %d, want method %d", params.Method, UMM_AL_QURA)
	}

	params = GetMethodParameters(DUBAI)
	if math.Abs(params.FajrAngle-18.2) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.2)
	}
	if math.Abs(params.IshaAngle-18.2) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 18.2)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != DUBAI {
		t.Errorf("error; got method = %d, want method %d", params.Method, DUBAI)
	}

	params = GetMethodParameters(MOON_SIGHTING_COMMITTEE)
	if math.Abs(params.FajrAngle-18) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.0)
	}
	if math.Abs(params.IshaAngle-18) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 18.0)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != MOON_SIGHTING_COMMITTEE {
		t.Errorf("error; got method = %d, want method %d", params.Method, MOON_SIGHTING_COMMITTEE)
	}

	params = GetMethodParameters(NORTH_AMERICA)
	if math.Abs(params.FajrAngle-15) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 15.0)
	}
	if math.Abs(params.IshaAngle-15) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 15.0)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != NORTH_AMERICA {
		t.Errorf("error; got method = %d, want method %d", params.Method, NORTH_AMERICA)
	}

	params = GetMethodParameters(KUWAIT)
	if math.Abs(params.FajrAngle-18) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.0)
	}
	if math.Abs(params.IshaAngle-17.5) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 17.5)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != KUWAIT {
		t.Errorf("error; got method = %d, want method %d", params.Method, KUWAIT)
	}

	params = GetMethodParameters(QATAR)
	if math.Abs(params.FajrAngle-18) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 18.0)
	}
	if math.Abs(params.IshaAngle-0) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 0.0)
	}
	if float64(params.IshaInterval) != 90 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 90)
	}
	if params.Method != QATAR {
		t.Errorf("error; got method = %d, want method %d", params.Method, QATAR)
	}

	params = GetMethodParameters(OTHER)
	if math.Abs(params.FajrAngle-0) > 0.000001 {
		t.Errorf("error; got fajr angle = %.2f; want fajr angle = %.2f", params.FajrAngle, 0.0)
	}
	if math.Abs(params.IshaAngle-0) > 0.000001 {
		t.Errorf("error; got isha angle = %.2f; want isha angle = %.2f", params.IshaAngle, 0.0)
	}
	if float64(params.IshaInterval) != 0 {
		t.Errorf("error; got isha interval = %d; want isha interval = %d", params.IshaInterval, 0)
	}
	if params.Method != OTHER {
		t.Errorf("error; got method = %d, want method %d", params.Method, OTHER)
	}
}
