package alert

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

type Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error)
	GetSingleAlert(token, orgId, envId, product, alertId string) (*alerts.AlertResponse, error)
	UpdateAlert(token, orgId, envId, product, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	DeleteSingleAlert(token, orgId, envId, product, alertId string) error
	GetAlertHistory(token, orgId, envId, product, alertId string) (*[]alerts.AlertHistoryResponse, error)
	GetResourceAlertHistory(token, orgId, envId, product, resourceId string) (*[]alerts.ResourceAlertHistory, error)
}
