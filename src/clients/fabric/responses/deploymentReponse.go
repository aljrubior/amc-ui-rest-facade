package responses

type DeploymentResponse struct {
	Id               string      `json:"id"`
	Name             string      `json:"name"`
	CreationDate     int64       `json:"creationDate"`
	LastModifiedDate int64       `json:"lastModifiedDate"`
	Target           Target      `json:"target"`
	Status           string      `json:"status"`
	Application      Application `json:"application"`
}
