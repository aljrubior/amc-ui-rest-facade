package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (t DefaultService) GetSingleAlert(token, orgId, envId, alertId string) (*alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetSingleAlert(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
