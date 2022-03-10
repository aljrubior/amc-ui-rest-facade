package responses

type IPAddress struct {
	RegionId   string `json:"regionId"`
	RegionName string `json:"regionName"`
	Type       string `json:"type"`
	Address    string `json:"address"`
	State      string `json:"state"`
}
