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
	applicationManagerV1Service ApplicationManagerV1Service,
	cloudhubService CloudhubService) DefaultAlertService {

	return DefaultAlertService{
		applicationManagerV1Service: applicationManagerV1Service,
		cloudhubService:             cloudhubService,
	}
}

type DefaultAlertService struct {
	applicationManagerV1Service ApplicationManagerV1Service
	cloudhubService             CloudhubService
}

func (service DefaultAlertService) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {
	var alerts []alerts.AlertResponse

	resp, err := service.applicationManagerV1Service.GetAlerts(token, orgId, envId)

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

	var alerts []alerts.AlertResponse

	if product == HYBRID_PRODUCT {
		resp, err := service.applicationManagerV1Service.PostAlert(token, orgId, envId, request)

		if err != nil {
			return nil, err
		}

		alerts = append(alerts, *resp...)

		return &alerts, err
	}

	if product == CLOUDHUB_PRODUCT {
		resp, err := service.cloudhubService.PostAlert(token, orgId, envId, request)

		if err != nil {
			return nil, err
		}

		alerts = append(alerts, *resp...)

		return &alerts, err
	}

	return &alerts, errors.New("//TODO: Implement this")
}

func (t DefaultAlertService) GetSingleAlert(token, orgId, envId, product, alertId string) (*[]alerts.AlertResponse, error) {

	var alerts []alerts.AlertResponse

	if product == HYBRID_PRODUCT {
		resp, err := t.applicationManagerV1Service.GetSingleAlert(token, orgId, envId, alertId)

		if err != nil {
			return nil, err
		}

		alerts = append(alerts, *resp...)

		return &alerts, err
	}

	if product == CLOUDHUB_PRODUCT {
		resp, err := t.cloudhubService.GetSingleAlert(token, orgId, envId, alertId)

		if err != nil {
			return nil, err
		}

		alerts = append(alerts, *resp...)

		return &alerts, err
	}

	return &alerts, errors.New("//TODO: Implement this")

}
