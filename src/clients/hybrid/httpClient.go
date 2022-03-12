package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

type HttpClient interface {
	GetAlerts(token, orgId, envId string) (*alerts.AlertsResponse, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*alerts.AlertsResponse, error)
	PatchSingleAlert(token, orgId, envId, alertId string, body []byte) (*alerts.AlertsResponse, error)
	PostAlert(token, orgId, envId string, body []byte) (*alerts.AlertsResponse, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*alerts.AlertHistoriesResponse, error)
	GetResourceAlertHistory(token, orgId, envId, resourceId string) (*alerts.ResourceAlertHistoriesResponse, error)

	GetApplications(token, orgId, envId string) (*responses.ApplicationsResponse, error)
}
