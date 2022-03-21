package model

type ServerArtifact struct {
	Target             ArtifactTarget `json:"target"`
	LastReportedStatus string         `json:"lastReportedStatus"`
	Message            string         `json:"message"`
}
