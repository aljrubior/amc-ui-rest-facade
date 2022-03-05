package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

type CloudhubService interface {
	GetAlerts(token string) (*[]alerts.AlertResponse, error)
}
