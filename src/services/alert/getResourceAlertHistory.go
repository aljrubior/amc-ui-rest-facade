package alert

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultService) GetResourceAlertHistory(token, orgId, envId, product, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	if product == HYBRID_PRODUCT {

		return t.hybridService.GetResourceAlertHistory(token, orgId, envId, resourceId)
	}

	if product == CLOUDHUB_PRODUCT {

		return nil, errors.New("//TODO: Not Implemented")
	}

	return nil, errors.New("//TODO: Implement this")
}
