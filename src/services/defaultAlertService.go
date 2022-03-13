package services

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

const (
	HYBRID_PRODUCT   = "hybrid"
	CLOUDHUB_PRODUCT = "cloudhub"
)

func NewDefaultAlertService(
	hybridService HybridService,
	cloudhubService CloudhubService) DefaultAlertService {

	return DefaultAlertService{
		hybridService:   hybridService,
		cloudhubService: cloudhubService,
	}
}

type DefaultAlertService struct {
	hybridService   HybridService
	cloudhubService CloudhubService
}

func (service DefaultAlertService) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {
	var alerts []alerts.AlertResponse

	resp, err := service.hybridService.GetAlerts(token, orgId, envId)

	if err != nil {
		// Log Error
	}

	alerts = append(alerts, *resp...)

	resp, err = service.cloudhubService.GetAlerts(token)

	if err != nil {
		// Log Error
	}

	alerts = append(alerts, *resp...)

	return &alerts, nil
}

func (service DefaultAlertService) PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return service.hybridService.PostAlert(token, orgId, envId, request)
	}

	if product == CLOUDHUB_PRODUCT {

		return service.cloudhubService.PostAlert(token, orgId, envId, request)
	}

	return nil, errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) GetSingleAlert(token, orgId, envId, product, alertId string) (*[]alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetSingleAlert(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.GetSingleAlert(token, orgId, envId, alertId)
	}

	return nil, errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) UpdateAlert(token, orgId, envId, product, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.UpdateSingleAlert(token, orgId, envId, alertId, request)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.UpdateAlert(token, orgId, envId, alertId, request)
	}

	return nil, errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) DeleteSingleAlert(token, orgId, envId, product, alertId string) error {

	if product == HYBRID_PRODUCT {

		return t.hybridService.DeleteSingleAlert(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.DeleteSingleAlert(token, orgId, envId, alertId)
	}

	return errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) GetAlertHistory(token, orgId, envId, product, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	var alerts []alerts.AlertHistoryResponse

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetAlertHistory(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.GetAlertHistory(token, orgId, envId, alertId)
	}

	return &alerts, errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) GetResourceAlertHistory(token, orgId, envId, product, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetResourceAlertHistory(token, orgId, envId, resourceId)
	}

	if product == CLOUDHUB_PRODUCT {

		return nil, errors.New("//TODO: Not Implemented")
	}

	return nil, errors.New("//TODO: Implement this")
}
