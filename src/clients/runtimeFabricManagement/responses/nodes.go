package responses

type Node struct {
	Id                        string `json:"id"`
	Location                  string `json:"location"`
	IsAvailableForDeployments bool   `json:"isAvailableForDeployments"`
	Availability              string `json:"availability"`
}
