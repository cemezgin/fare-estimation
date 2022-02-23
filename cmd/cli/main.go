package cli

import (
	"fare-estimation/internal/file"

	"github.com/spf13/cobra"
)

func Execute() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "execute",
		Short: "execute script for fare estimation",
		Run: func(cmd *cobra.Command, args []string) {
			file.WriteToCsv()
		},
	}

	return cmd
}
