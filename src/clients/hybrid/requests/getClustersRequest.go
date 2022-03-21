package requests

import (
	"fmt"
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
	"net/http"
)

func NewGetClustersRequest(
	config *config.HybridConfigClient,
	bearerToken,
	orgId,
	envId string) *GetClustersRequest {

	return &GetClustersRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetClustersRequest struct {
	clients.BaseHttpRequest
	config *config.HybridConfigClient
	bearerToken,
	orgId,
	envId string
}

func (t *GetClustersRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.ClustersPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetClustersRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddDefaultHeaders(req, t.orgId, t.envId, t.bearerToken)

	return req
}
