package responses

type DeploymentsResponse struct {
	Total int                  `json:"total"`
	Items []DeploymentResponse `json:"items"`
}
