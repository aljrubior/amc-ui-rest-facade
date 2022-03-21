package fabric

import "github.com/aljrubior/go-facade/clients/fabric/responses"

type Service interface {
	GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error)
}
