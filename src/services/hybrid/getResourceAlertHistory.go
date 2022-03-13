package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (t DefaultService) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	resp, err := t.httpClient.GetResourceAlertHistory(token, orgId, envId, resourceId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
