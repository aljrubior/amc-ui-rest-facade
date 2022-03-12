package target

type Response struct {
	Id               string   `json:"id"`
	Type             string   `json:"type"`
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	Nodes            []Node   `json:"nodes"`
	Provider         string   `json:"provider"`
	Defaults         Defaults `json:"defaults"`
	EnhancedSecurity bool     `json:"enhancedSecurity"`
}
