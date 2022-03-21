package hybrid

import (
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

func (t DefaultDatasource) CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.Response, error) {

	return t.hybridService.CreateAlert(token, orgId, envId, request)
}
