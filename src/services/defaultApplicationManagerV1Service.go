package services

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/applicationManagerV1"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
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

func (service DefaultApplicationManagerV1Service) PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

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

func (service DefaultApplicationManagerV1Service) GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error) {

	resp, err := service.httpClient.GetSingleAlert(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultApplicationManagerV1Service) UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

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

func (service DefaultApplicationManagerV1Service) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return service.httpClient.DeleteSingleAlert(token, orgId, envId, alertId)
}

func (service DefaultApplicationManagerV1Service) GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	resp, err := service.httpClient.GetAlertHistory(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (service DefaultApplicationManagerV1Service) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	resp, err := service.httpClient.GetResourceAlertHistory(token, orgId, envId, resourceId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
