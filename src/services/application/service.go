package application

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/application"

type Service interface {
	GetApplications(token, orgId, envId string) (*application.DataResponse, error)
}
