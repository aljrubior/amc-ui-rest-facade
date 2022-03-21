package alert

import (
	"errors"
	"fmt"
)

func (t DefaultService) DeleteSingleAlert(token, orgId, envId, product, alertId string) error {

	datasource, ok := t.datasources[product]

	if !ok {
		// TODO: Implement this
		return errors.New(fmt.Sprintf("Alerts not supported for product '%s'", product))
	}

	datasource.DeleteSingleAlert(token, orgId, envId, alertId)

	return errors.New("//TODO: Implement this")
}
