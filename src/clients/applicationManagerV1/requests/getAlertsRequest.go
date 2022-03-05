package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetAlertsRequest(
	config *config.ApplicationManagerV1ConfigClient,
	bearerToken,
	orgId,
	envId string) *GetAlertsRequest {

	return &GetAlertsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetAlertsRequest struct {
	clients.BaseHttpRequest
	config *config.ApplicationManagerV1ConfigClient
	bearerToken,
	orgId,
	envId string
}

func (request *GetAlertsRequest) buildUri() string {

	protocol := request.config.Protocol
	host := request.config.Host
	port := request.config.Port
	path := fmt.Sprintf("%s%s", request.config.Path, request.config.AlertsPath)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (request *GetAlertsRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, nil)

	request.AddDefaultHeaders(req, request.orgId, request.envId, request.bearerToken)

	return req
}
