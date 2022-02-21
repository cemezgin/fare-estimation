package calculate

import (
	"math"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

const radius = 6371 // Earth radius in km

func (origin Coordinates) Distance(destination Coordinates) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))

	b := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return radius * b
}

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
