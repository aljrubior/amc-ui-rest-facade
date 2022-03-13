package responses

type FeatureFlags struct {
	SupportsDefaultEndpoint bool `json:"supportsDefaultEndpoint"`
	SupportsPlaceHolder     bool `json:"supportsPlaceHolder"`
}
