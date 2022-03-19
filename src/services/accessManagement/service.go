package accessManagement

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/accessManagement/responses"
	"github.com/aljrubior/amc-ui-rest-facade/services/accessManagement/requests"
)

type Service interface {
	GetMembers(token, orgId, envId string) (*[]responses.MemberResponse, error)
	PostAuthorize(token, orgId, envId string, request requests.AuthorizeRequest) (*responses.AuthorizeResponse, error)
}
