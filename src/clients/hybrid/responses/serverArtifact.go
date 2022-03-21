package responses

type ServerArtifact struct {
	Id                 int      `json:"id"`
	TimeCreated        int64    `json:"timeCreated"`
	TimeUpdated        int64    `json:"timeUpdated"`
	ArtifactName       string   `json:"artifactName"`
	Artifact           Artifact `json:"artifact"`
	DeploymentId       int      `json:"deploymentId"`
	ServerId           int      `json:"serverId"`
	LastReportedStatus string   `json:"lastReportedStatus"`
	DesiredStatus      string   `json:"desiredStatus"`
	Message            string   `json:"message"`
	Discovered         bool     `json:"discovered"`
}
