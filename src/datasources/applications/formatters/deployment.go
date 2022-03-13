package formatters

type Deployment struct {
	Id                 string            `json:"id"`
	Artifact           Artifact          `json:"artifact"`
	LastReportedStatus string            `json:"lastReportedStatus"`
	Details            map[string]string `json:"details"`
	Target             Target            `json:"target"`
	Application        Application       `json:"application"`
}
