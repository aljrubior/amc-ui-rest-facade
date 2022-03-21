package alerts

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

type Datasource interface {
	CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.Response, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)
	GetAlerts(token, orgId, envId string) (*[]alerts.Response, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error)
	UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error)
	GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error)
}
