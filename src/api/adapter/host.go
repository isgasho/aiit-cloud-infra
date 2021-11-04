package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/mi-bear/infra-control/usecase/host"
)

type createHostRequestBody struct {
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}

func NewCreateHostInputPortFromRequest(r *http.Request) (*host.CreateHostInputPort, error) {
	var input createHostRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &host.CreateHostInputPort{
		Name:  input.Name,
		Limit: input.Limit,
	}, nil
}
