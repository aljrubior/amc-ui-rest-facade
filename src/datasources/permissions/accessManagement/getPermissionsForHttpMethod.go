package accessManagement

import (
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"github.com/aljrubior/go-facade/errors"
)

func (t DefaultDatasource) GetPermissionsForHttpMethod(token, orgId, envId, httpMethod string) *responses.AuthorizeResponse {

	result := make(responses.AuthorizeResponse, 0)

	resp, err := t.GetHybridPermissionForHttpMethod(token, orgId, envId, httpMethod)

	if _, ok := err.(errors.UnauthorizedError); !ok {
		// TODO: Log Error
	}

	result = append(result, *resp...)

	resp, err = t.GetRuntimeManagerPermissionForHttpMethod(token, orgId, envId, httpMethod)

	if _, ok := err.(errors.UnauthorizedError); !ok {
		// TODO: Log Error
	}

	result = append(result, *resp...)

	resp, err = t.GetTokenizationPermissionForHttpMethod(token, orgId, envId, httpMethod)

	if _, ok := err.(errors.UnauthorizedError); !ok {
		// TODO: Log Error
	}

	result = append(result, *resp...)

	return &result
}
