package alert

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultService) GetAlerts(token, orgId, envId string) (*[]alerts.Response, error) {
	var alerts []alerts.Response

	for _, v := range t.datasources {

		resp, err := v.GetAlerts(token, orgId, envId)

		if err != nil {
			// TODO: Log Error
		}

		alerts = append(alerts, *resp...)
	}

	return &alerts, nil
}
