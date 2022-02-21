package main

import (
	"fare-estimation/cmd/cli"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main()  {
	rootCmd := &cobra.Command{
		Use:   "fare-estimation",
		Short: "",
		Long:  "",
	}
	rootCmd.AddCommand(cli.Execute())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}