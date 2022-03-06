package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

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
