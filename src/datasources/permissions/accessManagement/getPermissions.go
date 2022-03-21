package accessManagement

import (
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"github.com/aljrubior/go-facade/datasources/permissions/accessManagement/formatters"
	"github.com/aljrubior/go-facade/datasources/permissions/accessManagement/transformers"
	"github.com/aljrubior/go-facade/model/responses/permission"
	"net/http"
)

func (t DefaultDatasource) GetPermissions(token, orgId, envId string) permission.Response {

	rawAuthorizations := make(map[string]*responses.AuthorizeResponse, 0)

	// GET permissions

	resp := t.GetPermissionsForHttpMethod(token, orgId, envId, http.MethodGet)

	rawAuthorizations[http.MethodGet] = resp

	// POST permissions

	resp = t.GetPermissionsForHttpMethod(token, orgId, envId, http.MethodPost)

	rawAuthorizations[http.MethodPost] = resp

	// PUT permissions

	resp = t.GetPermissionsForHttpMethod(token, orgId, envId, http.MethodPut)

	rawAuthorizations[http.MethodPut] = resp

	// PATCH permissions

	resp = t.GetPermissionsForHttpMethod(token, orgId, envId, http.MethodPatch)

	rawAuthorizations[http.MethodPatch] = resp

	// DELETE permissions

	resp = t.GetPermissionsForHttpMethod(token, orgId, envId, http.MethodDelete)

	rawAuthorizations[http.MethodDelete] = resp

	formattedAuthorizations := formatters.NewDefaultFormatter(&rawAuthorizations).Format()

	return transformers.NewDefaultTransformer(formattedAuthorizations).Transform()
}
