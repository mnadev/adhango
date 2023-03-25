package util

import "math"

func NormalizeWithBound(value float64, max float64) float64 {
	return value - (max * math.Floor(value/max))
}

func UnwindAngle(value float64) float64 {
	return NormalizeWithBound(value, 360.0)
}

func ClosestAngle(angle float64) float64 {
	if angle >= -180 && angle <= 180 {
		return angle

	}

	return angle - (360.0 * math.Round(angle/360.0))
}
