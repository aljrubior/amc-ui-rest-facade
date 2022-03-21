package hybrid

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultDatasource) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {

	return t.hybridService.GetResourceAlertHistory(token, orgId, envId, resourceId)
}
