package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (t DefaultDatasource) GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error) {

	return t.hybridService.GetSingleAlert(token, orgId, envId, alertId)
}
