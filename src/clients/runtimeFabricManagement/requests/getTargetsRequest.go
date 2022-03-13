package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetTargetsRequest(
	config *config.RuntimeFabricManagementClientConfig,
	bearerToken,
	orgId,
	envId string) *GetTargetsRequest {

	return &GetTargetsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetTargetsRequest struct {
	clients.BaseHttpRequest
	config *config.RuntimeFabricManagementClientConfig
	bearerToken,
	orgId,
	envId string
}

func (t *GetTargetsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.TargetsPath, t.orgId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetTargetsRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
