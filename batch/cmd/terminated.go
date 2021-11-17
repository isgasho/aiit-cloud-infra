package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// terminatedCmd represents the terminated command
var terminatedCmd = &cobra.Command{
	Use:   "terminated",
	Short: "terminate instance",
	Long: `terminate instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terminated called")
	},
}

func init() {
	rootCmd.AddCommand(terminatedCmd)
}
