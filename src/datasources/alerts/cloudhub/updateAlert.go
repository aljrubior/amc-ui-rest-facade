package cloudhub

import (
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

func (t DefaultDatasource) UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error) {

	return t.cloudhubService.UpdateAlert(token, orgId, envId, alertId, request)
}
