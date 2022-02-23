package cli

import (
	"fare-estimation/internal/file"
	"fare-estimation/pkg/calculate"
	"fmt"

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
					amount := calculate.Filter(rideList).FareAmount()
					c1 <- amount
				}()

				amountValue := <-c1
				fmt.Println(fmt.Sprintf("Ride ID: %d | Amount: %f", line[key].ID, amountValue))
			}
		},
	}

	return cmd
}
