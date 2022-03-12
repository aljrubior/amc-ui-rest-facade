package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric"
	"github.com/aljrubior/amc-ui-rest-facade/clients/fabric/responses"
)

func NewDefaultFabricService(httpClient fabric.HttpClient) DefaultFabricService {
	return DefaultFabricService{
		httpClient: httpClient,
	}
}

type DefaultFabricService struct {
	httpClient fabric.HttpClient
}

func (t DefaultFabricService) GetDeployments(token, orgId, envId string) (*[]responses.DeploymentResponse, error) {

	resp, err := t.httpClient.GetDeployments(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Items, nil
}
