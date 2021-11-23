package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var req = &InstanceCreateRequest{}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create instance",
	Long:  `create instance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := instanceCreate(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntVarP(&req.HostID, "host_id", "i", 0, "input create host ID")
	createCmd.Flags().StringVarP(&req.Name, "name", "n", "", "input instance name")
	createCmd.Flags().IntVarP(&req.Size, "size", "s", 0, "input instance size")
}

func instanceCreate(ir *InstanceCreateRequest) (interface{}, error) {
	jsonString, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, Endpoint+"/instances", bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
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
