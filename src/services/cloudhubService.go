package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

type CloudhubService interface {
	GetAlerts(token string) (*[]alerts.AlertResponse, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error)
	UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)
}
