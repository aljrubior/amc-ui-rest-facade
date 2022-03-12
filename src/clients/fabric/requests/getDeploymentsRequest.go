package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetDeploymentsRequest(
	config *config.FabricConfigClient,
	bearerToken,
	orgId,
	envId string) *GetDeploymentsRequest {

	return &GetDeploymentsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetDeploymentsRequest struct {
	clients.BaseHttpRequest
	config *config.FabricConfigClient
	bearerToken,
	orgId,
	envId string
}

func (t *GetDeploymentsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.DeploymentsPath, t.orgId, t.envId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetDeploymentsRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
