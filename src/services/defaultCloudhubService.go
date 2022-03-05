package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func NewDefaultCloudhubService(httpClient cloudhub.HttpClient) DefaultCloudhubService {

	return DefaultCloudhubService{
		httpClient: httpClient,
	}

}

type DefaultCloudhubService struct {
	httpClient cloudhub.HttpClient
}

func (service DefaultCloudhubService) GetAlerts(token string) (*[]alerts.AlertResponse, error) {

	resp, err := service.httpClient.GetAlerts(token)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
