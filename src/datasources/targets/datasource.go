package targets

import "github.com/aljrubior/go-facade/model/responses/target"

type Datasource interface {
	GetTargets(token, orgId, envId string) (*[]target.Response, error)
}
