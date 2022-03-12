package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/atlas/responses/target"

type AtlasService interface {
	GetTargets(token, orgId, envId, providerId string) (*[]target.Response, error)
}
