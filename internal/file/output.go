package file

import (
	"encoding/csv"
	"fare-estimation/pkg/calculate"
	"fmt"
	"os"
)
type RideAmountInterface interface {
	FareAmount() float64
}

func WriteToCsv()  {
	c1 := make(chan float64)
	csvList := [][]string{}
	csvFile, err := os.Create("assets/result.csv")
	if err != nil {
		fmt.Printf("failed creating file: %v", err)
		os.Exit(-1)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	records := ReadCsv("assets/paths.csv")
	for key, line := range records {
		rideList := line
		go func() {
			amount := RideAmountInterface.FareAmount(calculate.Filter(rideList))
			row := []string{fmt.Sprintf("%d", line[key].ID), fmt.Sprintf("%f", amount)}
			csvList = append(csvList, row)
			c1 <- amount
		}()

		amountValue := <-c1
		fmt.Println(fmt.Sprintf("RideCalculation ID: %d | Amount: %f ", line[key].ID, amountValue))

	}

	err = w.WriteAll(csvList)
	if err != nil {
		fmt.Printf("failed writing file: %v", err)
		os.Exit(-1)
	}
}
