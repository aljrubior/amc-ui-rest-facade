package requests

import (
	"fmt"
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
	"net/http"
)

func NewGetMembersRequest(
	config *config.AccessManagementConfigClient,
	bearerToken,
	orgId,
	envId string) *GetMembersRequest {

	return &GetMembersRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetMembersRequest struct {
	clients.BaseHttpRequest
	config *config.AccessManagementConfigClient
	bearerToken,
	orgId,
	envId string
}

func (t *GetMembersRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.MembersPath, t.orgId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetMembersRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	return req
}
