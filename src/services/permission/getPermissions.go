package permission

import (
	"github.com/aljrubior/amc-ui-rest-facade/model/responses/permission"
)

func (t DefaultService) GetPermissions(token, orgId, envId string) (*permission.DataResponse, error) {

	resp := t.datasource.GetPermissions(token, orgId, envId)

	return permission.NewDataResponse(resp), nil
}
