package permission

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/permission"

type Service interface {
	GetPermissions(token, orgId, envId string) (*permission.DataResponse, error)
}
