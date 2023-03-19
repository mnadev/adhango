package adhango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculationMethods(t *testing.T) {
	params := GetMethodParameters(MUSLIM_WORLD_LEAGUE)
	assert.InDelta(t, 18.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 17.0, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, MUSLIM_WORLD_LEAGUE, params.Method)

	params = GetMethodParameters(EGYPTIAN)
	assert.InDelta(t, 19.5, params.FajrAngle, 0.000001)
	assert.InDelta(t, 17.5, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, EGYPTIAN, params.Method)

	params = GetMethodParameters(KARACHI)
	assert.InDelta(t, 18.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 18.0, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, KARACHI, params.Method)

	params = GetMethodParameters(UMM_AL_QURA)
	assert.InDelta(t, 18.5, params.FajrAngle, 0.000001)
	assert.InDelta(t, 0, params.IshaAngle, 0.000001)
	assert.Equal(t, 90, params.IshaInterval)
	assert.Equal(t, UMM_AL_QURA, params.Method)

	params = GetMethodParameters(DUBAI)
	assert.InDelta(t, 18.2, params.FajrAngle, 0.000001)
	assert.InDelta(t, 18.2, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, DUBAI, params.Method)

	params = GetMethodParameters(MOON_SIGHTING_COMMITTEE)
	assert.InDelta(t, 18.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 18.0, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, MOON_SIGHTING_COMMITTEE, params.Method)

	params = GetMethodParameters(NORTH_AMERICA)
	assert.InDelta(t, 15.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 15.0, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, NORTH_AMERICA, params.Method)

	params = GetMethodParameters(KUWAIT)
	assert.InDelta(t, 18.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 17.5, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, KUWAIT, params.Method)

	params = GetMethodParameters(QATAR)
	assert.InDelta(t, 18.0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 0, params.IshaAngle, 0.000001)
	assert.Equal(t, 90, params.IshaInterval)
	assert.Equal(t, QATAR, params.Method)

	params = GetMethodParameters(OTHER)
	assert.InDelta(t, 0, params.FajrAngle, 0.000001)
	assert.InDelta(t, 0, params.IshaAngle, 0.000001)
	assert.Equal(t, 0, params.IshaInterval)
	assert.Equal(t, OTHER, params.Method)
}
