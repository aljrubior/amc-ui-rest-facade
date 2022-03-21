package cloudhub

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultDatasource) GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error) {

	return t.cloudhubService.GetSingleAlert(token, orgId, envId, alertId)
}
