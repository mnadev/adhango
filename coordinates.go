package adhango

import (
	"fmt"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinates(latitude float64, longitude float64) (*Coordinates, error) {
	if latitude > 90 || latitude < -90 {
		return nil, fmt.Errorf("latitude must be a number between -90 and 90 inclusive")
	}

	if longitude > 180 || longitude < -180 {
		return nil, fmt.Errorf("longitude must be a number between -180 and 180 inclusive")
	}

	return &Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
