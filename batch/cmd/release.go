package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release addresses",
	Long: `release addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release called")
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
