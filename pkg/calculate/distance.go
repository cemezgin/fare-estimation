package calculate

import (
	"math"
)

const radius = 6371 // Earth radius in km

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Distance struct {
}

func NewDistanceCalculator() *Distance {
	return &Distance{}
}

func (cd Distance) Distance(origin Coordinates, destination Coordinates) float64 {
	degreesLat := cd.degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := cd.degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(cd.degrees2radians(origin.Latitude))*
			math.Cos(cd.degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))

	b := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return radius * b
}

func (cd Distance) degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
