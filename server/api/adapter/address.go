package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mi-bear/infra-control/usecase/address"
)

type createAddressRequestBody struct {
	HostID     int    `json:"host_id"`
	IPAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
}

func NewCreateAddressInputPortFromRequest(r *http.Request) (*address.CreateAddressInputPort, error) {
	var input createAddressRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &address.CreateAddressInputPort{
		HostID:     input.HostID,
		IPAddress:  input.IPAddress,
		MacAddress: input.MacAddress,
	}, nil
}
