package cloudhub

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (t DefaultService) GetAlerts(token string) (*[]alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetAlerts(token)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
