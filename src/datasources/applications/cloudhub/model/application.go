package model

type Application struct {
	Id                  string      `json:"id"`
	Target              Target      `json:"target"`
	Artifact            Artifact    `json:"artifact"`
	MuleVersion         MuleVersion `json:"muleVersion"`
	IsDeploymentWaiting bool        `json:"isDeploymentWaiting"`
	LastReportedStatus  string      `json:"lastReportedStatus"`
	Details             Details     `json:"details"`
}
