package model

type Target struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Details     Details      `json:"details"`
	Deployments []Deployment `json:"deployments"`
	Status      string       `json:"status"`
}
