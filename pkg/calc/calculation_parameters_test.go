package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNightPortion(t *testing.T) {
	parameters := NewCalculationParametersBuilder().
		SetFajrAngle(18.0).
		SetIshaAngle(18.0).Build()
	parameters.HighLatitudeRule = MIDDLE_OF_THE_NIGHT
	np, err := parameters.NightPortions()
	assert.Nil(t, err)
	assert.InDelta(t, 0.5, np.Fajr, 0.001)
	assert.InDelta(t, 0.5, np.Isha, 0.001)

	parameters = NewCalculationParametersBuilder().
		SetFajrAngle(18.0).
		SetIshaAngle(18.0).Build()
	parameters.HighLatitudeRule = SEVENTH_OF_THE_NIGHT
	np, err = parameters.NightPortions()
	assert.Nil(t, err)
	assert.InDelta(t, 1.0/7.0, np.Fajr, 0.001)
	assert.InDelta(t, 1.0/7.0, np.Isha, 0.001)

	parameters = NewCalculationParametersBuilder().
		SetFajrAngle(10.0).
		SetIshaAngle(15.0).Build()
	parameters.HighLatitudeRule = TWILIGHT_ANGLE

	np, err = parameters.NightPortions()
	assert.Nil(t, err)
	assert.InDelta(t, 10.0/60.0, np.Fajr, 0.001)
	assert.InDelta(t, 15.0/60.0, np.Isha, 0.001)
}
