package permissions

import "github.com/aljrubior/go-facade/model/responses/permission"

type Datasource interface {
	GetPermissions(token, orgId, envId string) permission.Response
}
