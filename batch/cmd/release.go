package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release addresses",
	Long:  `release addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := getInstances(StateTerminated)
		if err != nil {
			fmt.Println(err)
		}

		for _, row := range res {
			instanceID := row.ID

			// Addresses.instance_id をnull にする API を呼び出す
			res, err := addressRelease(instanceID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
			fmt.Printf("Instance#%v terminated\n", instanceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}

func addressRelease(id int) (interface{}, error) {
	// TODO: implementation
	log.Printf("Instance#%v terminated\n", id)
	return nil, nil
}
