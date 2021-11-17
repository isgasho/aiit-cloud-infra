package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	
	"github.com/spf13/cobra"
)

var hostID int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete instance",
	Long: `delete instance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := instanceDeleteRequest(hostID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	
	deleteCmd.Flags().IntVarP(&hostID, "instance_id", "i", 0, "input delete instance ID")
}

func instanceDeleteRequest(id int) (interface{}, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/%v/%v", Endpoint, "instances", id), nil)
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
