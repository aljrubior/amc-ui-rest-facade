package responses

type TargetResponse struct {
	Id                        string       `json:"id"`
	Name                      string       `json:"name"`
	Version                   string       `json:"version"`
	Type                      string       `json:"type"`
	Runtimes                  []Runtime    `json:"runtimes"`
	Status                    string       `json:"status"`
	Environments              []string     `json:"environments"`
	IsAvailableForDeployments bool         `json:"isAvailableForDeployments"`
	Nodes                     []Node       `json:"nodes"`
	EnhancedSecurity          bool         `json:"enhancedSecurity"`
	FeatureFlags              FeatureFlags `json:"featureFlags"`
}
