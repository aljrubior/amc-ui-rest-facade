package permission

import "github.com/aljrubior/go-facade/model/responses/permission"

type Service interface {
	GetPermissions(token, orgId, envId string) (*permission.DataResponse, error)
}
