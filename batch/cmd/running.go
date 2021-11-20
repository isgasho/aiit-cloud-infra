package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:   "running",
	Short: "run instance",
	Long:  `run instance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := getInstances(StateStarting)
		if err != nil {
			fmt.Println(err)
		}

		for _, row := range res {
			instanceID := row.ID

			// Instance を Initializing にする
			if _, err := instanceStateUpdate(instanceID, StateInitializing); err != nil {
				fmt.Println(err)
			}

			// SSH Key を払い出す
			keys, err := NewKeys()
			if err != nil {
				fmt.Println(err)
			}
			privateKeyFilePath, err := keys.CreatePrivateKeyFile(instanceID)
			publicKeyFilePath, data, err := keys.CreatePublicKeyFile(instanceID)

			fmt.Printf("private: %v\n", privateKeyFilePath)
			fmt.Printf("public: %v\n", publicKeyFilePath)
			fmt.Println(data) //keys.data

			// TODO: keys.data を更新する

			// TODO: Instance を作る

			// TODO: Instance に Key を配置する

			// Instance を Running にする
			if _, err := instanceStateUpdate(instanceID, StateRunning); err != nil {
				fmt.Println(err)
			}

			// TODO: Private Key と設定情報を渡す

			fmt.Printf("Instance#%v running\n", instanceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(runningCmd)
}
