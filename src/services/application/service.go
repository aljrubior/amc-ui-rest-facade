package application

import "github.com/aljrubior/go-facade/model/responses/application"

type Service interface {
	GetApplications(token, orgId, envId string) (*application.DataResponse, error)
}
