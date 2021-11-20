package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// terminatedCmd represents the terminated command
var terminatedCmd = &cobra.Command{
	Use:   "terminated",
	Short: "terminate instance",
	Long:  `terminate instance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := getInstances(StateTerminating)
		if err != nil {
			fmt.Println(err)
		}

		for _, row := range res {
			instanceID := row.ID

			// TODO: Instance を物理的に削除する

			// Instance を Terminated にする
			res, err := instanceStateUpdate(instanceID, StateTerminated)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
			fmt.Printf("Instance#%v terminated\n", instanceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(terminatedCmd)
}
