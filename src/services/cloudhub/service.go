package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

type Service interface {
	GetAlerts(token string) (*[]alerts.AlertResponse, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error)
	UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)

	GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error)
}
