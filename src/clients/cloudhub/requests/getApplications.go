package requests

import (
	"fmt"
	"github.com/aljrubior/amc-ui-rest-facade/clients"
	"github.com/aljrubior/amc-ui-rest-facade/config"
	"net/http"
)

func NewGetApplications(
	config *config.CloudhubConfigClient,
	bearerToken string) *GetApplications {

	return &GetApplications{
		config:      config,
		bearerToken: bearerToken,
	}
}

type GetApplications struct {
	clients.BaseHttpRequest
	config      *config.CloudhubConfigClient
	bearerToken string
}

func (t *GetApplications) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.ApplicationsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetApplications) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddAuthorizationHeader(req, t.bearerToken)

	return req
}
