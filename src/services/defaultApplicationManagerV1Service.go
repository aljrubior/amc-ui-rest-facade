package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/applicationManagerV1"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func NewDefaultApplicationManagerV1Service(
	httpClient applicationManagerV1.HttpClient) DefaultApplicationManagerV1Service {
	return DefaultApplicationManagerV1Service{
		httpClient: httpClient,
	}
}

type DefaultApplicationManagerV1Service struct {
	httpClient applicationManagerV1.HttpClient
}

func (service DefaultApplicationManagerV1Service) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {

	resp, err := service.httpClient.GetAlerts(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
