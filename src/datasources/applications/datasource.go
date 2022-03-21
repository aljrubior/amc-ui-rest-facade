package applications

import "github.com/aljrubior/go-facade/model/responses/application"

type Datasource interface {
	GetApplications(token, orgId, envId string) (*[]application.Response, error)
}
