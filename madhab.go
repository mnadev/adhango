package adhango

type AsrJuristicMethod int64

const (
	SHAFI_HANBALI_MALIKI AsrJuristicMethod = iota

	HANAFI
)

var MadhabToShadowLengthMap = map[AsrJuristicMethod]ShadowLength{
	SHAFI_HANBALI_MALIKI: SINGLE,
	HANAFI:               DOUBLE,
}
