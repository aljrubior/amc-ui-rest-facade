package services

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

type AlertService interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error)
	PostAlert(token, orgId, envId, product string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
}
