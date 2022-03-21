package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetAlertsRequest(
	config *config.CloudhubConfigClient,
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
	config *config.CloudhubConfigClient
	bearerToken,
	orgId,
	envId string
}

func (t *GetAlertsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.AlertsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetAlertsRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddAuthorizationHeader(req, t.bearerToken)

	return req
}
