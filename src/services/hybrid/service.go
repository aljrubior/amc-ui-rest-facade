package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/server"
	"github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/serverGroup"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

type Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error)
	UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)
	GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error)

	GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error)
	GetServers(token, orgId, envId string) (*[]server.Response, error)
	GetServerGroups(token, orgId, envId string) (*[]serverGroup.Response, error)
	GetClusters(token, orgId, envId string) (*[]cluster.Response, error)
}
