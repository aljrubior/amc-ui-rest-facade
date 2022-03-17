package alert

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

func (t DefaultService) UpdateAlert(token, orgId, envId, product, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.UpdateSingleAlert(token, orgId, envId, alertId, request)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.UpdateAlert(token, orgId, envId, alertId, request)
	}

	return nil, errors.New("//TODO: Implement this")
}
