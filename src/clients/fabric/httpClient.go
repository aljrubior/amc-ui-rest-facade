package fabric

import "github.com/aljrubior/go-facade/clients/fabric/responses"

type HttpClient interface {
	GetDeployments(token, orgId, envId string) (*responses.DeploymentsResponse, error)
}
