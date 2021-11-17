package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:   "running",
	Short: "run instance",
	Long: `run instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running called")
	},
}

func init() {
	rootCmd.AddCommand(runningCmd)
}
