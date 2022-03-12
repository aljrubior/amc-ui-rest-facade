package services

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

func NewDefaultHybridService(
	httpClient hybrid.HttpClient) DefaultHybridService {
	return DefaultHybridService{
		httpClient: httpClient,
	}
}

type DefaultHybridService struct {
	httpClient hybrid.HttpClient
}

func (service DefaultHybridService) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {

	resp, err := service.httpClient.GetAlerts(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := service.httpClient.PostAlert(token, orgId, envId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error) {

	resp, err := service.httpClient.GetSingleAlert(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := service.httpClient.PatchSingleAlert(token, orgId, envId, alertId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return service.httpClient.DeleteSingleAlert(token, orgId, envId, alertId)
}

func (service DefaultHybridService) GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	resp, err := service.httpClient.GetAlertHistory(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	resp, err := service.httpClient.GetResourceAlertHistory(token, orgId, envId, resourceId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultHybridService) GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error) {

	resp, err := service.httpClient.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
