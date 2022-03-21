package target

type RuntimeVersion struct {
	BaseVersion string `json:"baseVersion"`
	Tag         string `json:"tag"`
	MinimumTag  string `json:"minimumTag"`
}
