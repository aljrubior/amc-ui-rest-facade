package applications

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/application"

type Datasource interface {
	GetApplications(token, orgId, envId string) (*[]application.Response, error)
}
