package presenter

import (
	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/address"
)

type simpleAddressView struct {
	ID         int    `json:"id"`
	IPAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
}

func mapAddressToSimpleView(address *model.Address) *simpleAddressView {
	return &simpleAddressView{
		ID:         address.ID,
		IPAddress:  address.IPAddress,
		MacAddress: address.MacAddress,
	}
}

type createAddressResponse struct {
	Address *simpleAddressView `json:"address"`
}

func NewCreateAddressPresenter(output *address.CreateAddressOutputPort) *createAddressResponse {
	return &createAddressResponse{mapAddressToSimpleView(output.Address)}
}
