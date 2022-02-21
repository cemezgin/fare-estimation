package cli

import (
	"fare-estimation/internal/file"
	"fare-estimation/internal/ride"
	"fare-estimation/pkg/calculate"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Execute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "execute script for fare estimation",
		Run: func(cmd *cobra.Command, args []string) {
			c1 := make(chan float64)

			records := file.ReadCsv("assets/paths.csv")
			for key, line := range records {
				rideList := line
				go func() {
					amount := filter(rideList).FareAmount()
					c1 <- amount
				}()

				amountValue := <-c1
				fmt.Println(fmt.Sprintf("Ride ID: %d | Amount: %f", line[key].ID, amountValue))
			}
		},
	}

	return cmd
}

func filter(line []ride.Ride) *calculate.Ride {
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

		pointA := calculate.Coordinates{lat, long}
		pointB := calculate.Coordinates{lat2, long2}

		distanceCalculate := pointA.Distance(pointB)
		subTime := line[key+1].Time.Sub(r.Time)

		timeDistance := subTime.Hours()
		speedHourly := distanceCalculate / timeDistance
		if speedHourly < 100 {
			//@todo find
		}

		tuples = append(tuples, ride.RideTuples{
			distanceCalculate,
			timeDistance,
			speedHourly,
		})

	}

	return calculate.NewRide(tuples, line[0].Time)
}
