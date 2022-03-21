package accessManagement

import (
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"github.com/aljrubior/go-facade/datasources/permissions/accessManagement/model"
	"github.com/aljrubior/go-facade/services/accessManagement/requests"
)

func (t DefaultDatasource) GetHybridPermissionForHttpMethod(token, orgId, envId, httpMethod string) (*responses.AuthorizeResponse, error) {

	resourceNames := []string{model.APPLICATIONS_RESOURCE_NAME, model.SERVERS_RESOURCE_NAME, model.SERVER_GROUPS_RESOURCE_NAME, model.CLUSTERS_RESOURCE_NAME}

	request := requests.NewAuthorizeRequest(orgId, envId, model.HYBRID_NAMESPACE, httpMethod, resourceNames)

	return t.accessManagementService.PostAuthorize(token, orgId, envId, request)
}
