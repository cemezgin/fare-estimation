package calculate

import (
	"fare-estimation/internal/ride"
	"fmt"
	"time"
)

const FARE_AMOUNT_PER_KM_DAY = 0.74
const FARE_AMOUNT_PER_KM_NIGHT = 1.30
const FARE_AMOUNT_IDLE = 11.90
const MIN_FARE_AMOUNT = 3.47
const FLAG_AMOUNT = 1.30

type RideCalculation struct {
	rideList  []ride.RideTuples
	firstTime time.Time
}

type DistanceCalculatorInterface interface {
	Distance(origin Coordinates, destination Coordinates) float64
}

func NewRide(rideList []ride.RideTuples, firstTime time.Time) *RideCalculation {
	return &RideCalculation{firstTime: firstTime, rideList: rideList}
}

func (r RideCalculation) FareAmount() float64 {
	var distance float64
	var idleTime float64
	flag := FLAG_AMOUNT

	for _, fare := range r.rideList {
		if fare.SpeedHourly > 10 {
			distance = fare.TimeDistance + distance
		} else {
			idleTime = fare.TimeDistance + idleTime
		}
	}

	totalAmount := r.calculateMoving(distance) + r.calculateIdle(idleTime) + flag

	if totalAmount < MIN_FARE_AMOUNT {
		return MIN_FARE_AMOUNT
	}

	return totalAmount
}

func (r RideCalculation) calculateMoving(totalDistance float64) float64 {
	zero := "2006-01-02T00:00:00Z"
	five := "2006-01-02T05:00:00Z"
	var totalAmount float64

	zeroTime, err := time.Parse(time.RFC3339, zero)
	if err != nil {
		fmt.Println(err)
	}
	fiveTime, err := time.Parse(time.RFC3339, five)
	if err != nil {
		fmt.Println(err)
	}

	if r.firstTime.Hour() > zeroTime.Hour() && r.firstTime.Hour() < fiveTime.Hour() {
		totalAmount = (totalDistance * FARE_AMOUNT_PER_KM_NIGHT)
	} else {
		totalAmount = (totalDistance * FARE_AMOUNT_PER_KM_DAY)
	}

	return totalAmount
}

func (r RideCalculation) calculateIdle(totalTime float64) float64 {
	return FARE_AMOUNT_IDLE * totalTime
}

func Filter(line []ride.Ride) *RideCalculation {
	tuples := []ride.RideTuples{}
	for key, r := range line {
		if key+1 == len(line) {
			continue
		}

		distanceObject := NewDistanceCalculator()
		pointA := Coordinates{r.Lat, r.Long}
		pointB := Coordinates{line[key+1].Lat, line[key+1].Long}
		distanceCalculate := DistanceCalculatorInterface.Distance(distanceObject, pointA, pointB)

		subTime := line[key+1].Time.Sub(r.Time)
		timeDistance := subTime.Hours()
		speedHourly := distanceCalculate / timeDistance

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
