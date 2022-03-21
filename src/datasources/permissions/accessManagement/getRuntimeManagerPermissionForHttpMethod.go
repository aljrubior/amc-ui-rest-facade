package accessManagement

import (
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"
	"github.com/aljrubior/go-facade/services/accessManagement/requests"
)

func (t DefaultDatasource) GetRuntimeManagerPermissionForHttpMethod(token, orgId, envId, httpMethod string) (*responses.AuthorizeResponse, error) {

	resourceNames := []string{model.ALERTS_RESOURCE_NAME}

	request := requests.NewAuthorizeRequest(orgId, envId, model.RUNTIME_MANAGER_NAMESPACE, httpMethod, resourceNames)

	return t.accessManagementService.PostAuthorize(token, orgId, envId, request)
}
