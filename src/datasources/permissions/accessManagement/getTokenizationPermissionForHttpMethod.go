package accessManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
	"github.com/aljrubior/amc-ui-rest-facade/datasources/permissions/accessManagement/model"
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement/requests"
)

func (t DefaultDatasource) GetTokenizationPermissionForHttpMethod(token, orgId, envId, httpMethod string) (*responses.AuthorizeResponse, error) {

	resourceNames := []string{model.SERVICES_NAMESPACE}

	request := requests.NewAuthorizeRequest(orgId, envId, model.TOKENIZATION_NAMESPACE, httpMethod, resourceNames)

	return t.accessManagementService.PostAuthorize(token, orgId, envId, request)
}
