package targets

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/target"

type TargetDatasource interface {
	GetTargets(token, orgId, envId string) (*[]target.Response, error)
}