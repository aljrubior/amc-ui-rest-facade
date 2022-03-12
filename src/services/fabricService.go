package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"

type FabricService interface {
	GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error)
}
