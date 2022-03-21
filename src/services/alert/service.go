package alert

import (
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

type Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.Response, error)
	GetSingleAlert(token, orgId, envId, product, alertId string) (*alerts.Response, error)
	UpdateAlert(token, orgId, envId, product, alertId string, request requests.AlertRequest) (*[]alerts.Response, error)
	PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.Response, error)
	DeleteSingleAlert(token, orgId, envId, product, alertId string) error
	GetAlertHistory(token, orgId, envId, product, alertId string) (*[]alerts.AlertHistoryResponse, error)
	GetResourceAlertHistory(token, orgId, envId, product, resourceId string) (*[]alerts.ResourceAlertHistory, error)
}
