package responses

type Runtime struct {
	Type    string           `json:"type"`
	Version []RuntimeVersion `json:"versions"`
}
