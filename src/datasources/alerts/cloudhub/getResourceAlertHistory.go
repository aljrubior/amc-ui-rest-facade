package cloudhub

import "github.com/aljrubior/go-facade/clients/responses/alerts"

func (t DefaultDatasource) GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error) {
	// NOT SUPPORTED
	return &[]alerts.ResourceAlertHistory{}, nil
}
