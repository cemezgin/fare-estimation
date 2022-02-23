package cli

import (
	"encoding/csv"
	"fare-estimation/internal/file"
	"fare-estimation/pkg/calculate"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type RideAmountInterface interface {
	FareAmount() float64
}

func Execute() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "execute script for fare estimation",
		Run: func(cmd *cobra.Command, args []string) {
			c1 := make(chan float64)
			csvList := [][]string{}
			csvFile, err := os.Create("assets/result.csv")
			if err != nil {
				fmt.Println("failed creating file: %s", err)
				os.Exit(-1)
			}
			defer csvFile.Close()

			w := csv.NewWriter(csvFile)
			defer w.Flush()

			records := file.ReadCsv("assets/paths.csv")
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
				fmt.Println("failed writing file: %s", err)
				os.Exit(-1)
			}
		},
	}

	return cmd
}
