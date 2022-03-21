package alert

import (
	"errors"
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

func (t DefaultService) GetSingleAlert(token, orgId, envId, product, alertId string) (*alerts.Response, error) {

	datasource, ok := t.datasources[product]

	if !ok {
		// TODO: Implement this
		return nil, errors.New(fmt.Sprintf("Alerts not supported for product '%s'", product))
	}

	return datasource.GetSingleAlert(token, orgId, envId, alertId)
}
