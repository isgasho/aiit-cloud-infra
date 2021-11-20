package cmd

type InstanceCreateRequest struct {
	HostID int    `json:"host_id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
}
