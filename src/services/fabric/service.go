package fabric

import "github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"

type Service interface {
	GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error)
}
