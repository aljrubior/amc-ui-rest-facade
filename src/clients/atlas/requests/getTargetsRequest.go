package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetTargetsRequest(
	config *config.AtlasConfigClient,
	bearerToken,
	orgId,
	envId,
	providerId string) *GetTargetsRequest {

	return &GetTargetsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
		providerId:  providerId,
	}
}

type GetTargetsRequest struct {
	clients.BaseHttpRequest
	config *config.AtlasConfigClient
	bearerToken,
	orgId,
	envId,
	providerId string
}

func (t *GetTargetsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := fmt.Sprintf(t.config.TargetsPath, t.orgId, t.providerId)

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetTargetsRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	return req
}
