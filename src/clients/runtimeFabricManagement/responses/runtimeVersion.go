package responses

type RuntimeVersion struct {
	BaseVersion string `json:"baseVersion"`
	Tag         string `json:"tag"`
	ReleaseDate int64  `json:"releaseDate"`
	MinimumTag  string `json:"minimumTag"`
}
