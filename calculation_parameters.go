package adhango

import "fmt"

type CalculationParameters struct {
	//  The method used to do the calculation
	Method CalculationMethod

	// The angle of the sun used to calculate fajr
	FajrAngle float64

	// The angle of the sun used to calculate isha
	IshaAngle float64

	// Minutes after Maghrib (if set, the time for Isha will be Maghrib plus IshaInterval)
	IshaInterval int

	// The juristic method used to calculate Asr
	Madhab AsrJuristicMethod

	// Rules for placing bounds on Fajr and Isha for high latitude areas
	HighLatitudeRule HighLatitudeRule

	// Used to optionally add or subtract a set amount of time from each prayer time
	Adjustments PrayerAdjustments

	// Used for method adjustments
	MethodAdjustments PrayerAdjustments
}

type CalculationParametersBuilder struct {
	//  The method used to do the calculation
	Method CalculationMethod

	// The angle of the sun used to calculate fajr
	FajrAngle float64

	// The angle of the sun used to calculate isha
	IshaAngle float64

	// Minutes after Maghrib (if set, the time for Isha will be Maghrib plus IshaInterval)
	IshaInterval int

	// The juristic method used to calculate Asr
	Madhab AsrJuristicMethod

	// Rules for placing bounds on Fajr and Isha for high latitude areas
	HighLatitudeRule HighLatitudeRule

	// Used to optionally add or subtract a set amount of time from each prayer time
	Adjustments PrayerAdjustments

	// Used for method adjustments
	MethodAdjustments PrayerAdjustments
}

func NewCalculationParametersBuilder() *CalculationParametersBuilder {
	return &CalculationParametersBuilder{
		Method:            OTHER,
		FajrAngle:         0.0,
		IshaAngle:         0.0,
		IshaInterval:      0,
		Madhab:            SHAFI_HANBALI_MALIKI,
		HighLatitudeRule:  MIDDLE_OF_THE_NIGHT,
		Adjustments:       PrayerAdjustments{},
		MethodAdjustments: PrayerAdjustments{},
	}
}

func (cpb *CalculationParametersBuilder) SetMethod(m CalculationMethod) *CalculationParametersBuilder {
	cpb.Method = m
	return cpb
}

func (cpb *CalculationParametersBuilder) SetFajrAngle(fajrAngle float64) *CalculationParametersBuilder {
	cpb.FajrAngle = fajrAngle
	return cpb
}

func (cpb *CalculationParametersBuilder) SetIshaAngle(ishaAngle float64) *CalculationParametersBuilder {
	cpb.IshaAngle = ishaAngle
	return cpb
}

func (cpb *CalculationParametersBuilder) SetIshaInterval(ishaInterval int) *CalculationParametersBuilder {
	cpb.IshaInterval = ishaInterval
	return cpb
}

func (cpb *CalculationParametersBuilder) SetMadhab(madhab AsrJuristicMethod) *CalculationParametersBuilder {
	cpb.Madhab = madhab
	return cpb
}

func (cpb *CalculationParametersBuilder) SetHighLatitudeRule(highLatitudeRule HighLatitudeRule) *CalculationParametersBuilder {
	cpb.HighLatitudeRule = highLatitudeRule
	return cpb
}

func (cpb *CalculationParametersBuilder) SetAdjustments(adjustments PrayerAdjustments) *CalculationParametersBuilder {
	cpb.Adjustments = adjustments
	return cpb
}

func (cpb *CalculationParametersBuilder) SetMethodAdjustments(methodAdjustments PrayerAdjustments) *CalculationParametersBuilder {
	cpb.MethodAdjustments = methodAdjustments
	return cpb
}

func (cpb *CalculationParametersBuilder) Build() *CalculationParameters {
	return &CalculationParameters{
		Method:            cpb.Method,
		FajrAngle:         cpb.FajrAngle,
		IshaAngle:         cpb.IshaAngle,
		IshaInterval:      cpb.IshaInterval,
		Madhab:            cpb.Madhab,
		HighLatitudeRule:  cpb.HighLatitudeRule,
		Adjustments:       cpb.Adjustments,
		MethodAdjustments: cpb.MethodAdjustments,
	}
}

func (c *CalculationParameters) NightPortions() (*NightPortions, error) {
	if c.HighLatitudeRule == MIDDLE_OF_THE_NIGHT {
		return NewNightPortions(1.0/2.0, 1.0/2.0)
	}

	if c.HighLatitudeRule == SEVENTH_OF_THE_NIGHT {
		return NewNightPortions(1.0/7.0, 1.0/7.0)
	}

	if c.HighLatitudeRule == TWILIGHT_ANGLE {
		return NewNightPortions(c.FajrAngle/60.0, c.IshaAngle/60.0)
	}

	return nil, fmt.Errorf("invalid high latitude rule")
}
