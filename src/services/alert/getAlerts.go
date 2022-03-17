package alert

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (service DefaultService) GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error) {
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
