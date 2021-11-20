package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release addresses",
	Long:  `release addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := getTerminatedInstances()
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

func getTerminatedInstances() ([]InstanceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", Endpoint, "instances"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("state", StateTerminated)
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

func addressRelease(id int) (interface{}, error) {
	// TODO: implementation
	return nil, nil
}
