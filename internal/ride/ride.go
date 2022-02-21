package ride

import "time"

type Ride struct {
	ID   int
	Lat  string
	Long string
	Time time.Time
}

type RideTuples struct {
	DistanceCalculate float64
	TimeDistance float64
	SpeedHourly float64
}
