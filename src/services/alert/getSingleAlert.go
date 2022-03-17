package alert

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultService) GetSingleAlert(token, orgId, envId, product, alertId string) (*alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetSingleAlert(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.GetSingleAlert(token, orgId, envId, alertId)
	}

	return nil, errors.New("//TODO: Implement this")
}
