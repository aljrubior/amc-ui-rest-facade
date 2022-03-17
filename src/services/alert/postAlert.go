package alert

import (
	"errors"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

func (service DefaultService) PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	if product == HYBRID_PRODUCT {

		return service.hybridService.CreateAlert(token, orgId, envId, request)
	}

	if product == CLOUDHUB_PRODUCT {

		return service.cloudhubService.CreateAlert(token, orgId, envId, request)
	}

	return nil, errors.New("//TODO: Implement this")
}
