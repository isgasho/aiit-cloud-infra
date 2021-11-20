package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// terminatedCmd represents the terminated command
var terminatedCmd = &cobra.Command{
	Use:   "terminated",
	Short: "terminate instance",
	Long:  `terminate instance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := getTerminatingInstances()
		if err != nil {
			fmt.Println(err)
		}

		for _, row := range res {
			instanceID := row.ID

			// TODO: Instance を物理的に削除する

			// Instance を Terminated にする
			res, err := instanceUpdate(instanceID)
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

func getTerminatingInstances() ([]InstanceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", Endpoint, "instances"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("state", StateTerminating)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("file close failed")
		}
	}(resp.Body)

	var result []InstanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func instanceUpdate(id int) (interface{}, error) {
	req, err := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("%v/instances/%v/state/%v", Endpoint, id, StateTerminated), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("file close failed")
		}
	}(resp.Body)

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
