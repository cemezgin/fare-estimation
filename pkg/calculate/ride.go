package calculate

import (
	"fare-estimation/internal/ride"
	"fmt"
	"strconv"
	"time"
)

const FARE_AMOUNT_PER_KM_DAY = 0.74
const FARE_AMOUNT_PER_KM_NIGHT = 1.30
const FARE_AMOUNT_IDLE = 11.90

type Ride struct {
	rideList  []ride.RideTuples
	firstTime time.Time
}

type DistanceCalculatorInterface interface {
	Distance(origin Coordinates, destination Coordinates) float64
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

func Filter(line []ride.Ride) *Ride {
	tuples := []ride.RideTuples{}
	for key, r := range line {
		if key+1 == len(line) {
			continue
		}

		lat, err := strconv.ParseFloat(r.Lat, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		long, err := strconv.ParseFloat(r.Long, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		lat2, err := strconv.ParseFloat(line[key+1].Lat, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		long2, err := strconv.ParseFloat(line[key+1].Long, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		distanceObject := NewDistanceCalculator()
		pointA := Coordinates{lat, long}
		pointB := Coordinates{lat2, long2}
		distanceCalculate := DistanceCalculatorInterface.Distance(distanceObject,pointA,pointB)

		subTime := line[key+1].Time.Sub(r.Time)
		timeDistance := subTime.Hours()
		speedHourly := distanceCalculate / timeDistance

		fmt.Println(r.Lat, r.Long, r.Time)
		fmt.Println(line[key+1].Lat, line[key+1].Long, line[key+1].Time)
		fmt.Println(distanceCalculate, timeDistance, speedHourly)
		if speedHourly > 100 {
			//@todo remove from list
		}

		tuples = append(tuples, ride.RideTuples{
			distanceCalculate,
			timeDistance,
			speedHourly,
		})
	}

	return NewRide(tuples, line[0].Time)
}
