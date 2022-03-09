package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

type CloudhubService interface {
	GetAlerts(token string) (*[]alerts.AlertResponse, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error)
	PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
}
