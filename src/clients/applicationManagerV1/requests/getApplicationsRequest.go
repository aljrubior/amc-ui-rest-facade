package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetApplicationsRequest(
	config *config.ApplicationManagerV1ConfigClient,
	bearerToken,
	orgId,
	envId string) *GetApplicationsRequest {

	return &GetApplicationsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetApplicationsRequest struct {
	clients.BaseHttpRequest
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId string
}

func (request *GetApplicationsRequest) buildUri() string {

	protocol := request.config.Protocol
	host := request.config.Host
	port := request.config.Port
	path := request.config.ApplicationsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetApplicationsRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
