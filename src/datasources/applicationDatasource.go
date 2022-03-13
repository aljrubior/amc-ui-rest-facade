package datasources

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/application"

type ApplicationDatasource interface {
	GetApplications(token, orgId, envId string) (*[]application.Response, error)
}
