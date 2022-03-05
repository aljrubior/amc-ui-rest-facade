package applicationManagerV1

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
)

type HttpClient interface {
	GetAlerts(token, orgId, envId string) (*alerts.AlertsResponse, error)
}
