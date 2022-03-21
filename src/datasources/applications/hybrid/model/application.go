package model

type Application struct {
	Id                 int      `json:"id"`
	Target             Target   `json:"target"`
	Artifact           Artifact `json:"artifact"`
	LastReportedStatus string   `json:"lastReportedStatus"`
	DesiredStatus      string   `json:"desiredStatus"`
	Details            Details  `json:"details"`
	Uptime             int      `json:"uptime"`
}
