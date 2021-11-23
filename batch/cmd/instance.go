package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getInstances(state string) ([]InstanceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", Endpoint, "instances"), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("state", state)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("file close failed")
		}
	}(resp.Body)

	var result InstancesResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Instances, nil
}

func instanceStateUpdate(id int, state string) (interface{}, error) {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%v/instances/%v/state/%v", Endpoint, id, state), nil)
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
