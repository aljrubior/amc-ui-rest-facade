package services

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/server"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/serverGroup"
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

func (t DefaultHybridService) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetAlerts(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PostAlert(token, orgId, envId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetSingleAlert(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PatchSingleAlert(token, orgId, envId, alertId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	return t.httpClient.DeleteSingleAlert(token, orgId, envId, alertId)
}

func (t DefaultHybridService) GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	resp, err := t.httpClient.GetAlertHistory(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	resp, err := t.httpClient.GetResourceAlertHistory(token, orgId, envId, resourceId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error) {

	resp, err := t.httpClient.GetApplications(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetServerGroups(token, orgId, envId string) (*[]serverGroup.Response, error) {

	resp, err := t.httpClient.GetServerGroups(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetClusters(token, orgId, envId string) (*[]cluster.Response, error) {

	resp, err := t.httpClient.GetClusters(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultHybridService) GetServers(token, orgId, envId string) (*[]server.Response, error) {

	resp, err := t.httpClient.GetServers(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
