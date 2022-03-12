package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/atlas"
	"github.com/aljrubior/amc-ui-rest-facade/clients/atlas/responses/target"
)

func NewDefaultAtlasService(httpClient atlas.HttpClient) DefaultAtlasService {
	return DefaultAtlasService{
		httpClient: httpClient,
	}
}

type DefaultAtlasService struct {
	httpClient atlas.HttpClient
}

func (t DefaultAtlasService) GetTargets(token, orgId, envId, providerId string) (*[]target.Response, error) {

	return t.httpClient.GetTargets(token, orgId, envId, providerId)
}
