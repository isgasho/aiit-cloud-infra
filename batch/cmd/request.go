package cmd

type InstanceCreateRequest struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}
