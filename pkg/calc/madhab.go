package calc

import (
	util "github.com/mnadev/adhango/pkg/util"
)

type AsrJuristicMethod int64

const (
	SHAFI_HANBALI_MALIKI AsrJuristicMethod = iota

	HANAFI
)

var MadhabToShadowLengthMap = map[AsrJuristicMethod]util.ShadowLength{
	SHAFI_HANBALI_MALIKI: util.SINGLE,
	HANAFI:               util.DOUBLE,
}
