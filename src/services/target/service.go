package target

import "github.com/aljrubior/go-facade/model/responses/target"

type Service interface {
	GetTargets(token, orgId, envId string) (*target.DataResponse, error)
}
