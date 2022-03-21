package alert

import (
	"errors"
	"fmt"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

func (t DefaultService) PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.Response, error) {

	datasource, ok := t.datasources[product]

	if !ok {
		// TODO: Implement this
		return nil, errors.New(fmt.Sprintf("Alerts not supported for product '%s'", product))
	}

	return datasource.CreateAlert(token, orgId, envId, request)
}
