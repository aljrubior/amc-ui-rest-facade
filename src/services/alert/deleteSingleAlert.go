package alert

import "errors"

func (t DefaultService) DeleteSingleAlert(token, orgId, envId, product, alertId string) error {

	if product == HYBRID_PRODUCT {

		return t.hybridService.DeleteSingleAlert(token, orgId, envId, alertId)
	}

	if product == CLOUDHUB_PRODUCT {

		return t.cloudhubService.DeleteSingleAlert(token, orgId, envId, alertId)
	}

	return errors.New("//TODO: Implement this")
}
