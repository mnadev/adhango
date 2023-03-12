package adhango

type NightPortions struct {
	Fajr float64
	Isha float64
}

func NewNightPortions(fajr float64, isha float64) (*NightPortions, error) {
	return &NightPortions{fajr, isha}, nil
}
