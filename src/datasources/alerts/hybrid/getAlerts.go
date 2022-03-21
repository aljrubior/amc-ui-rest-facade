package hybrid

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultDatasource) GetAlerts(token, orgId, envId string) (*[]alerts.Response, error) {

	return t.hybridService.GetAlerts(token, orgId, envId)
}
