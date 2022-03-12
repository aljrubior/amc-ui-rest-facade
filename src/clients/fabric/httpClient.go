package fabric

import "github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"

type HttpClient interface {
	GetDeployments(token, orgId, envId string) (*responses.DeploymentsResponse, error)
}
