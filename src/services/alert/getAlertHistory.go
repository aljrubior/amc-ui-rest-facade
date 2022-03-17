package alert

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultService) GetAlertHistory(token, orgId, envId, product, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	var alerts []alerts.AlertHistoryResponse

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetAlertHistory(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.GetAlertHistory(token, orgId, envId, alertId)
	}

	return &alerts, errors.New("//TODO: Implement this")
}
