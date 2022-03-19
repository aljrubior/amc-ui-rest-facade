package permissions

import "github.com/aljrubior/amc-ui-rest-facade/model/responses/permission"

type Datasource interface {
	GetPermissions(token, orgId, envId string) permission.Response
}
