package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/instance"
)

type createInstanceRequestBody struct {
	HostID int    `json:"host_id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
}

func NewCreateInstanceInputPortFromRequest(r *http.Request) (*instance.CreateInstanceInputPort, error) {
	var input createInstanceRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &instance.CreateInstanceInputPort{
		HostID: input.HostID,
		Name:   input.Name,
		Size:   input.Size,
		Key:    nil,
	}, nil
}

func NewGetInstanceInputPortFromRequest(r *http.Request) (*instance.GetInstanceInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &instance.GetInstanceInputPort{
		ID: ID,
	}, nil
}

func NewUpdateInstanceInputPortFromRequest(r *http.Request) (*instance.UpdateInstanceInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	// TODO: State validation
	// 1: Starting
	// 2: Initializing
	// 3: Running
	// 4: Terminating
	// 5: Terminated
	state, err := strconv.Atoi(vars["state"])
	if err != nil {
		return nil, err
	}

	return &instance.UpdateInstanceInputPort{
		ID:    ID,
		State: model.State(state),
	}, nil
}

func NewDeleteInstanceInputPortFromRequest(r *http.Request) (*instance.DeleteInstanceInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &instance.DeleteInstanceInputPort{
		ID: ID,
	}, nil
}

func NewListInstancesInputPortFromRequest(r *http.Request) (*instance.ListInstancesInputPort, error) {
	q := r.URL.Query()
	hostID, err := strconv.Atoi(q.Get("host_id"))
	if err != nil {
		return nil, err
	}

	return &instance.ListInstancesInputPort{
		HostID: hostID,
	}, nil
}
