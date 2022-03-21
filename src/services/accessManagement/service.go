package accessManagement

import (
	"github.com/aljrubior/go-facade/clients/accessManagement/responses"
	"github.com/aljrubior/go-facade/services/accessManagement/requests"
)

type Service interface {
	GetMembers(token, orgId, envId string) (*[]responses.MemberResponse, error)
	PostAuthorize(token, orgId, envId string, request requests.AuthorizeRequest) (*responses.AuthorizeResponse, error)
}
