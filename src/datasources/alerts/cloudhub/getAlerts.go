package cloudhub

import (
	"github.com/aljrubior/go-facade/clients/responses/alerts"
)

func (t DefaultDatasource) GetAlerts(token, orgId, envId string) (*[]alerts.Response, error) {

	return t.cloudhubService.GetAlerts(token, orgId, envId)
}
