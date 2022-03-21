package atlas

import "github.com/aljrubior/go-facade/clients/atlas/responses/target"

type HttpClient interface {
	GetTargets(token, orgId, envId, providerId string) (*[]target.Response, error)
}
