package services

import "github.com/aljrubior/amc-ui-rest-facade/clients/applicationManagerV1"

type DefaultAlertService struct {
	applicationManagerV1.HttpClient
}
