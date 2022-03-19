package accessManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
)

type HttpClient interface {
	GetMembers(token, orgId, envId string) (*responses.DataResponse, error)
	PostAuthorize(token, orgId, envId string, body []byte) (*responses.AuthorizeResponse, error)
}
