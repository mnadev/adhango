package adhango

import (
	"math"
	"testing"
)

func TestNormalizeWithBound(t *testing.T) {
	testCases := []struct {
		value     float64
		max       float64
		want      float64
		tolerance float64
	}{
		{2.0, -5, -3, 1e-5},
		{-4.0, -5.0, -4, 1e-5},
		{-6.0, -5.0, -1, 1e-5},
		{-1.0, 24, 23, 1e-5},
		{1.0, 24.0, 1, 1e-5},
		{49.0, 24, 1, 1e-5},
		{361.0, 360, 1, 1e-5},
		{360.0, 360, 0, 1e-5},
		{259.0, 360, 259, 1e-5},
		{2592.0, 360, 72, 1e-5},
		{360.1, 360, 0.1, 1e-2},
	}
	for _, tc := range testCases {
		got := NormalizeWithBound(tc.value, tc.max)
		if math.Abs(got-tc.want) > tc.tolerance {
			t.Errorf("error; got = %.2f wanted = %.2f", got, tc.want)
		}
	}
}

func TestUnwindAngle(t *testing.T) {
	testCases := []struct {
		value float64
		want  float64
	}{
		{-45.0, 315},
		{361.0, 1},
		{360.0, 0},
		{259.0, 259},
		{2592.0, 72},
	}
	for _, tc := range testCases {
		got := UnwindAngle(tc.value)
		if math.Abs(got-tc.want) > 1e-5 {
			t.Errorf("error; got = %.2f wanted = %.2f", got, tc.want)
		}
	}
}

func TestClosestAngle(t *testing.T) {
	testCases := []struct {
		angle     float64
		want      float64
		tolerance float64
	}{
		{360.0, 0, 1e-6},
		{361.0, 1, 1e-6},
		{1.0, 1, 1e-6},
		{-1.0, -1, 1e-6},
		{-181.0, 179, 1e-6},
		{180.0, 180, 1e-6},
		{359.0, -1, 1e-6},
		{-359.0, 1, 1e-6},
		{1261.0, -179, 1e-6},
		{-360.1, -0.1, 1e-2},
	}
	for _, tc := range testCases {
		got := ClosestAngle(tc.angle)
		if math.Abs(got-tc.want) > tc.tolerance {
			t.Errorf("error; got = %.2f wanted = %.2f", got, tc.want)
		}
	}
}
