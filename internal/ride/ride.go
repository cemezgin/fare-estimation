package ride

import "time"

type Ride struct {
	ID   int
	Lat  float64
	Long float64
	Time time.Time
}

type RideTuples struct {
	DistanceCalculate float64
	TimeDistance      float64
	SpeedHourly       float64
}
