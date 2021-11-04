package presenter

import (
	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/instance"
)

type simpleInstanceView struct {
	ID         int         `json:"id"`
	HostID     int         `json:"host_id"`
	Name       string      `json:"name"`
	State      model.State `json:"state"`
	Size       int         `json:"size"`
	IPAddress  string      `json:"ip_address"`
	MacAddress string      `json:"mac_address"`
	Key        string      `json:"key"`
}

func mapInstanceToSimpleView(instance *model.Instance) *simpleInstanceView {
	return &simpleInstanceView{
		ID:         instance.ID,
		HostID:     instance.HostID,
		Name:       instance.Name,
		State:      instance.State,
		Size:       instance.Size,
		IPAddress:  instance.Address.IPAddress,
		MacAddress: instance.Address.MacAddress,
		Key:        instance.Key.Data,
	}
}

type createInstanceResponse struct {
	Instance *simpleInstanceView `json:"instance"`
}

func NewCreateInstancePresenter(output *instance.CreateInstanceOutputPort) *createInstanceResponse {
	return &createInstanceResponse{mapInstanceToSimpleView(output.Instance)}
}

type getInstanceResponse struct {
	Instance *simpleInstanceView `json:"instance"`
}

func NewGetInstancePresenter(output *instance.GetInstanceOutputPort) *getInstanceResponse {
	return &getInstanceResponse{mapInstanceToSimpleView(output.Instance)}
}

type updateInstanceResponse struct {
	Instance *simpleInstanceView `json:"instance"`
}

func NewUpdateInstancePresenter(output *instance.UpdateInstanceOutputPort) *updateInstanceResponse {
	return &updateInstanceResponse{mapInstanceToSimpleView(output.Instance)}
}

type deleteInstanceResponse struct {
	Instance *simpleInstanceView `json:"instance"`
}

func NewDeleteInstancePresenter(output *instance.DeleteInstanceOutputPort) *deleteInstanceResponse {
	return &deleteInstanceResponse{mapInstanceToSimpleView(output.Instance)}
}

type listInstanceResponse struct {
	Instances []*simpleInstanceView `json:"instances"`
}

func NewListInstancesPresenter(output *instance.ListInstancesOutputPort) *listInstanceResponse {
	views := make([]*simpleInstanceView, len(output.Instances))

	for i, row := range output.Instances {
		views[i] = mapInstanceToSimpleView(row)
	}
	return &listInstanceResponse{views}
}
