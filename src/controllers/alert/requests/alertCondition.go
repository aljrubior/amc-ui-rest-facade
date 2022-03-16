package requests

type AlertCondition struct {
	Resources    []string `json:"resources"`
	Type         string   `json:"type"`
	ResourceType string   `json:"resourceType"`
}
