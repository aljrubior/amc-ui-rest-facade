package target

type Runtime struct {
	Type                      string           `json:"type"`
	Versions                  []RuntimeVersion `json:"versions"`
	Environments              []string         `json:"environments"`
	IsAvailableForDeployments bool             `json:"isAvailableForDeployments"`
}
