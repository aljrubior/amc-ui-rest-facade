package cloudhub

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

type HttpClient interface {
	GetAlerts(token string) (*alerts.AlertsResponse, error)
}
