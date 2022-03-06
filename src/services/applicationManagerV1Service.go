package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

type ApplicationManagerV1Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.AlertResponse, error)
}
