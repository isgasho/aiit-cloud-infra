package presenter

import (
	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/host"
)

type simpleHostView struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}

func mapHostToSimpleView(host *model.Host) *simpleHostView {
	return &simpleHostView{
		ID:    host.ID,
		Name:  host.Name,
		Limit: host.Limit,
	}
}

type createHostResponse struct {
	Host *simpleHostView `json:"host"`
}

func NewCreateHostPresenter(output *host.CreateHostOutputPort) *createHostResponse {
	return &createHostResponse{mapHostToSimpleView(output.Host)}
}
