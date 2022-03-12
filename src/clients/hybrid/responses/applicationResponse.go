package responses

type ApplicationResponse struct {
	Id                 int              `json:"id"`
	TimeCreated        int64            `json:"timeCreated"`
	TimeUpdated        int64            `json:"timeUpdated"`
	Name               string           `json:"name"`
	Uptime             int              `json:"uptime"`
	DesiredStatus      string           `json:"desiredStatus"`
	LastReportedStatus string           `json:"lastReportedStatus"`
	Started            bool             `json:"started"`
	ServerArtifacts    []ServerArtifact `json:"serverArtifacts"`
	Artifact           Artifact         `json:"artifact"`
	Target             Target           `json:"target"`
}
