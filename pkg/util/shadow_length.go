package util

type ShadowLength int64

const (
	SINGLE ShadowLength = iota

	DOUBLE
)

var ShadowLengthToFloatMap = map[ShadowLength]float64{
	SINGLE: 1.0,
	DOUBLE: 2.0,
}
