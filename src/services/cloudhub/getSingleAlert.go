package cloudhub

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultService) GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error) {

	return t.httpClient.GetSingleAlert(token, orgId, envId, alertId)
}
