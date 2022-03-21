package hybrid

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultService) GetAlerts(token, orgId, envId string) (*[]alerts.Response, error) {

	resp, err := t.httpClient.GetAlerts(token, orgId, envId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
