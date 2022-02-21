package calculate

import (
	"fare-estimation/internal/ride"
	"fmt"
	"time"
)

const FARE_AMOUNT_PER_KM_DAY = 0.74
const FARE_AMOUNT_PER_KM_NIGHT = 1.30
const FARE_AMOUNT_IDLE = 11.90

type Ride struct {
	rideList  []ride.RideTuples
	firstTime time.Time
}

func NewRide(rideList []ride.RideTuples, firstTime time.Time) *Ride {
	return &Ride{firstTime: firstTime, rideList: rideList}
}

func (r Ride) FareAmount() float64 {
	var distance float64
	var idleTime float64

	for _, fare := range r.rideList {
		if fare.SpeedHourly > 10 {
			distance = fare.TimeDistance + distance
		} else {
			idleTime = fare.TimeDistance + idleTime
		}
	}

	return r.calculateMoving(distance) + r.calculateIdle(idleTime)

}

func (r Ride) calculateMoving(totalDistance float64) float64 {
	zero := "2006-01-02T00:00:00Z"
	five := "2006-01-02T05:00:00Z"

	zeroTime, err := time.Parse(time.RFC3339, zero)
	if err != nil {
		fmt.Println(err)
	}
	fiveTime, err := time.Parse(time.RFC3339, five)
	if err != nil {
		fmt.Println(err)
	}

	if r.firstTime.Hour() > zeroTime.Hour() && r.firstTime.Hour() < fiveTime.Hour() {
		totalDistance = totalDistance * FARE_AMOUNT_PER_KM_NIGHT
	} else {
		totalDistance = totalDistance * FARE_AMOUNT_PER_KM_DAY
	}

	return totalDistance
}

func (r Ride) calculateIdle(totalTime float64) float64 {
	return FARE_AMOUNT_IDLE / totalTime
}
