package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultDatasource) GetAlerts(token, orgId, envId string) (*[]alerts.Response, error) {

	return t.cloudhubService.GetAlerts(token, orgId, envId)
}
