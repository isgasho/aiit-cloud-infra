package adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/instance"
)

type createInstanceRequestBody struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func NewCreateInstanceInputPortFromRequest(r *http.Request) (*instance.CreateInstanceInputPort, error) {
	var input createInstanceRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &instance.CreateInstanceInputPort{
		Name: input.Name,
		Size: input.Size,
		Key:  nil,
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

	state, err := model.StringToState(vars["state"])
	if err != nil {
		return nil, err
	}

	return &instance.UpdateInstanceInputPort{
		ID:    ID,
		State: state,
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
	state, err := model.StringToState(q.Get("state"))
	if err != nil {
		log.Println("StringToState parse error")
		return nil, err
	}

	return &instance.ListInstancesInputPort{
		State: state,
	}, nil
}
