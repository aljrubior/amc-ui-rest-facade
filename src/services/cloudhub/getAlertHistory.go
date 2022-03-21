package cloudhub

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultService) GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	resp, err := t.httpClient.GetAlertHistory(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
