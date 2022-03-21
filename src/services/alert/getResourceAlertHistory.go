package alert

import (
	"errors"
	"fmt"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
)

func (t DefaultService) GetResourceAlertHistory(token, orgId, envId, product, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	datasource, ok := t.datasources[product]

	if !ok {
		// TODO: Implement this
		return nil, errors.New(fmt.Sprintf("Alerts not supported for product '%s'", product))
	}

	return datasource.GetResourceAlertHistory(token, orgId, envId, resourceId)
}
