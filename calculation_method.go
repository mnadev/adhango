package adhango

type CalculationMethod int64

const (
	OTHER CalculationMethod = iota
	// Muslim World League
	// Uses Fajr angle of 18 and an Isha angle of 17
	MUSLIM_WORLD_LEAGUE
	// Egyptian General Authority of Survey
	// Uses Fajr angle of 19.5 and an Isha angle of 17.5
	EGYPTIAN
	// University of Islamic Sciences, Karachi
	// Uses Fajr angle of 18 and an Isha angle of 18
	KARACHI
	// Umm al-Qura University, Makkah
	// Uses a Fajr angle of 18.5 and an Isha angle of 90. Note: You should add a +30 minute custom
	// adjustment of Isha during Ramadan.
	UMM_AL_QURA
	// The Gulf Region
	// Uses Fajr and Isha angles of 18.2 degrees.
	DUBAI
	// Moonsighting Committee
	// Uses a Fajr angle of 18 and an Isha angle of 18. Also uses seasonal adjustment values.
	MOON_SIGHTING_COMMITTEE
	// Referred to as the ISNA method
	// This method is included for completeness, but is not recommended.
	// Uses a Fajr angle of 15 and an Isha angle of 15.
	NORTH_AMERICA
	// Kuwait
	// Uses a Fajr angle of 18 and an Isha angle of 17.5
	KUWAIT
	// Qatar
	// Modified version of Umm al-Qura that uses a Fajr angle of 18.
	QATAR
	// Singapore
	// Uses a Fajr angle of 20 and an Isha angle of 18
	SINGAPORE
	// UOIF
	// Uses a Fajr angle of 12 and an Isha angle of 12
	UOIF
)

func GetMethodParameters(method CalculationMethod) *CalculationParameters {
	cpb := NewCalculationParametersBuilder().SetMethod(method)
	switch method {
	case MUSLIM_WORLD_LEAGUE:
		cpb.SetFajrAngle(18.0).
			SetIshaAngle(17.0).
			SetMethodAdjustments(PrayerAdjustments{DhuhrAdj: 1})
	case EGYPTIAN:
		cpb.SetFajrAngle(19.5).
			SetIshaAngle(17.5).
			SetMethodAdjustments(PrayerAdjustments{DhuhrAdj: 1})
	case KARACHI:
		cpb.SetFajrAngle(18.0).
			SetIshaAngle(18.0).
			SetMethodAdjustments(PrayerAdjustments{DhuhrAdj: 1})
	case UMM_AL_QURA:
		cpb.SetFajrAngle(18.5).
			SetIshaInterval(90)
	case DUBAI:
		cpb.SetFajrAngle(18.2).
			SetIshaAngle(18.2).
			SetMethodAdjustments(PrayerAdjustments{
				SunriseAdj: -3,
				DhuhrAdj:   3,
				AsrAdj:     3,
				MaghribAdj: 3,
			})
	case MOON_SIGHTING_COMMITTEE:
		cpb.SetFajrAngle(18.0).
			SetIshaAngle(18.0).
			SetMethodAdjustments(PrayerAdjustments{
				DhuhrAdj:   5,
				MaghribAdj: 3,
			})
	case NORTH_AMERICA:
		cpb.SetFajrAngle(15.0).
			SetIshaAngle(15.0).
			SetMethodAdjustments(PrayerAdjustments{
				DhuhrAdj: 1,
			})
	case KUWAIT:
		cpb.SetFajrAngle(18.0).
			SetIshaAngle(17.5)
	case QATAR:
		cpb.SetFajrAngle(18.0).
			SetIshaInterval(90)
	case SINGAPORE:
		cpb.SetFajrAngle(20.0).
			SetIshaAngle(18.0).
			SetMethodAdjustments(PrayerAdjustments{
				DhuhrAdj: 1,
			})
	case UOIF:
		cpb.SetFajrAngle(12.0).
			SetIshaAngle(12.0)
	}

	return cpb.Build()
}
