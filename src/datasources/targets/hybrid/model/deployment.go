package model

type Deployment struct {
	Id                 int      `json:"id"`
	Artifact           Artifact `json:"artifact"`
	LastReportedStatus string   `json:"lastReportedStatus"`
}
