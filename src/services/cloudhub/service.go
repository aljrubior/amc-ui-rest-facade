package cloudhub

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

type Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.Response, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error)
	UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error)
	CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.Response, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)

	GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error)
}
