package cmd

type InstanceResponse struct {
	ID         int    `json:"id"`
	HostID     int    `json:"host_id"`
	Name       string `json:"name"`
	State      int    `json:"state"`
	Size       int    `json:"size"`
	IPAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
	Key        string `json:"key"`
}

type InstancesResponse struct {
	Instances []InstanceResponse `json:"instances"`
}
